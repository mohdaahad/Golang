package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func Authenticate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq LoginRequest
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
			return
		}

		var user User
		if err := db.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid username or password"})
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			Role:     user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "token": tokenString})
	}
}
