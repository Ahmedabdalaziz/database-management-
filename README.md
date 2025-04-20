# Golang MySQL User Management

This is a simple Go application that connects to a MySQL database and provides basic user management operations such as:

- Insert a user
- Update user details
- Delete a user
- Get a single user by ID
- Get all users
- Find users by city
- Count total users

## 🛠️ Requirements

- Go 1.18 or later
- MySQL server running locally
- `github.com/go-sql-driver/mysql` driver

## 🗃️ Database Setup

Make sure you have a database named `golang` and a table `users` with the following structure:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    city VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
