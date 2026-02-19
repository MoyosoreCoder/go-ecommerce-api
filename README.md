# E-commerce api 

## Description
RESTFUL backend API implementation for an e-commerce  web application built with Go. Features include user authentication, product management, and database management

## Features
- User authentication:
- Login user with jwt authentication
- Product:
- Add products (admin-only)
- List all products
- Optional if there is time for me: view, update, delete products
- PostgreSQL database for persistent storage
- RESTful API endpoints using Gorilla

## Tech Stack
- **Backend:** Go (Golang)
- **Routing:** Gorilla Mux
- **Database:** PostgreSQL
- **Authentication:** JWT

## Project Structure
go-ecommerce-api/
- README.md
- go.mod
- main.go
- models
- handlers
- routes

## Getting started
1. Install Go: https://golang.org/dl/
2. Install PostgreSQL: https://www.postgresql.org/download/  NOTE: enable the pgAdmin during installation
3. Clone this repository:
```bash
git clone https://github.com/MoyosoreCoder/go-ecommerce-api.git
```
4. Set up the project
- go mod init (your-github-project-link)
5. Run this command in the terminal
```bash
go run main.go
```