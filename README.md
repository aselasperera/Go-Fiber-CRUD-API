# Go Fiber CRUD API

This is a simple RESTful API built with [Go Fiber](https://gofiber.io) for handling CRUD operations on user resources. The API uses SQLite as the database and GORM as the ORM (Object-Relational Mapper).

## Features

- **Create** a new user
- **Read** existing users
- **Update** an existing user
- **Delete** a user

## Getting Started

### Prerequisites

- Go (version 1.16+ is recommended)
- Git (optional)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/aselasperera/go-fiber-crud-api.git
cd go-fiber-crud-api

2. Install dependencies
go mod tidy

3. Run the application
go run main.go  The server will start on http://localhost:3000.