-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- Create the roles table
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

-- Create the permissions table
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES roles(id),
    resource VARCHAR(50) NOT NULL,
    can_read BOOLEAN NOT NULL,
    can_write BOOLEAN NOT NULL
);

-- Create the user_roles table
CREATE TABLE IF NOT EXISTS user_roles (
    user_id INTEGER REFERENCES users(id),
    role_id INTEGER REFERENCES roles(id)
);

-- Create the customers table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT
);

-- Create the billing table
CREATE TABLE IF NOT EXISTS billing (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id),
    amount DECIMAL(10, 2) NOT NULL,
    billing_date DATE NOT NULL,
    status VARCHAR(50)
);

-- Create the payroll table
CREATE TABLE IF NOT EXISTS payroll (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER REFERENCES employees(id),
    salary DECIMAL(10, 2) NOT NULL,
    pay_date DATE NOT NULL,
    status VARCHAR(50)
);

-- Create the employees table
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    department VARCHAR(100),
    position VARCHAR(100)
);

-- Create the user_logs table
CREATE TABLE IF NOT EXISTS user_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    action VARCHAR(100),
    action_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert roles
INSERT INTO roles (role_name) VALUES ('Sales'), ('Accountant'), ('HR'), ('Administrator');

-- Insert permissions
INSERT INTO permissions (role_id, resource, can_read, can_write) VALUES
((SELECT id FROM roles WHERE role_name = 'Sales'), 'Customer Management', true, true),
((SELECT id FROM roles WHERE role_name = 'Sales'), 'Billing Management', true, true),
((SELECT id FROM roles WHERE role_name = 'Accountant'), 'Billing Management', true, false),
((SELECT id FROM roles WHERE role_name = 'Accountant'), 'Payroll Management', true, false),
((SELECT id FROM roles WHERE role_name = 'HR'), 'Payroll Management', true, true),
((SELECT id FROM roles WHERE role_name = 'Administrator'), 'User Management', true, true);

-- Insert users (example user data with hashed passwords)
INSERT INTO users (username, password) VALUES 
('sales_user', 'hashed_password1'),
('accountant_user', 'hashed_password2'),
('hr_user', 'hashed_password3'),
('admin_user', 'hashed_password4');

-- Map users to roles
INSERT INTO user_roles (user_id, role_id) VALUES
((SELECT id FROM users WHERE username = 'sales_user'), (SELECT id FROM roles WHERE role_name = 'Sales')),
((SELECT id FROM users WHERE username = 'accountant_user'), (SELECT id FROM roles WHERE role_name = 'Accountant')),
((SELECT id FROM users WHERE username = 'hr_user'), (SELECT id FROM roles WHERE role_name = 'HR')),
((SELECT id FROM users WHERE username = 'admin_user'), (SELECT id FROM roles WHERE role_name = 'Administrator'));

-- Insert sample customer data
INSERT INTO customers (name, email, phone, address) VALUES
('Customer One', 'customer1@example.com', '1234567890', '123 First St, City, Country'),
('Customer Two', 'customer2@example.com', '0987654321', '456 Second St, City, Country');

-- Insert sample billing data
INSERT INTO billing (customer_id, amount, billing_date, status) VALUES
((SELECT id FROM customers WHERE name = 'Customer One'), 100.00, '2024-01-01', 'Paid'),
((SELECT id FROM customers WHERE name = 'Customer Two'), 200.00, '2024-01-02', 'Pending');

-- Insert sample employee data
INSERT INTO employees (user_id, first_name, last_name, department, position) VALUES
((SELECT id FROM users WHERE username = 'sales_user'), 'Sales', 'User', 'Sales', 'Sales Representative'),
((SELECT id FROM users WHERE username = 'accountant_user'), 'Accountant', 'User', 'Accounting', 'Accountant'),
((SELECT id FROM users WHERE username = 'hr_user'), 'HR', 'User', 'HR', 'HR Manager'),
((SELECT id FROM users WHERE username = 'admin_user'), 'Admin', 'User', 'Admin', 'Administrator');
