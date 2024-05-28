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
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

The server will start on `http://localhost:3000`.

## API Endpoints

### Get All Users

- **URL:** `/users`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    [
      {
        "ID": 1,
        "Name": "John Doe",
        "Email": "john@example.com"
      },
      ...
    ]
    ```

### Get a Single User

- **URL:** `/users/:id`
- **Method:** `GET`
- **URL Params:**
  - `id=[integer]`
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "ID": 1,
      "Name": "John Doe",
      "Email": "john@example.com"
    }
    ```

### Create a New User

- **URL:** `/users`
- **Method:** `POST`
- **Data Params:**
  - `name=[string]`
  - `email=[string]`
- **Success Response:**
  - **Code:** 201
  - **Content:**
    ```json
    {
      "ID": 1,
      "Name": "John Doe",
      "Email": "john@example.com"
    }
    ```

### Update an Existing User

- **URL:** `/users/:id`
- **Method:** `PUT`
- **URL Params:**
  - `id=[integer]`
- **Data Params:**
  - `name=[string]`
  - `email=[string]`
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "ID": 1,
      "Name": "John Doe",
      "Email": "john@example.com"
    }
    ```

### Delete a User

- **URL:** `/users/:id`
- **Method:** `DELETE`
- **URL Params:**
  - `id=[integer]`
- **Success Response:**
  - **Code:** 200

## Project Structure

- `main.go`: The entry point of the application.
- `routes/user.go`: Defines the user-related routes.
- `models/user.go`: Defines the User model.
- `handlers/user.go`: Contains the handler functions for CRUD operations.
- `database/db.go`: Sets up and initializes the database.

## Acknowledgements
Go Fiber
GORM
SQLite

This updated `README.md` file includes the "Install dependencies" step formatted in the same style as the "Clone the repository" step, making it clear and consistent.

