
# Go Gin CRUD API

This is a simple CRUD API built with the Go Gin framework. The API manages user data with an in-memory database, allowing for basic user operations such as Create, Read, Update, and Delete (CRUD).

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Create User](#create-user)
  - [Get All Users](#get-all-users)
  - [Get User by ID](#get-user-by-id)
  - [Update User](#update-user)
  - [Delete User](#delete-user)

## Features
- Create, Read, Update, and Delete users.
- In-memory database for fast operations.
- RESTful API built with the Go Gin framework.

## Installation

1. **Clone the repository:**

   ```bash
   git clone github.com/gitish/polyglot_training.git
   cd golang/go-gin-crud
   ```

2. **Install dependencies:**

   Make sure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run golang/go-gin-crud/main.go
   ```

   The server will start on `http://localhost:8080`.

## Usage

You can use [Postman](https://www.postman.com/) or any HTTP client to interact with the API. Below are the available endpoints:

## API Endpoints

### Create User
- **Endpoint:** `POST /users`
- **Description:** Creates a new user.
- **Request Body:**
  ```json
  {
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```
- **Response:**
  ```json
  {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```

### Get All Users
- **Endpoint:** `GET /users`
- **Description:** Retrieves a list of all users.
- **Response:**
  ```json
  [
      {
          "id": 1,
          "name": "John Doe",
          "email": "johndoe@example.com"
      }
  ]
  ```

### Get User by ID
- **Endpoint:** `GET /users/:id`
- **Description:** Retrieves a user by their ID.
- **Example URL:** `GET /users/1`
- **Response:**
  ```json
  {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```

### Update User
- **Endpoint:** `PUT /users/:id`
- **Description:** Updates an existing user by ID.
- **Example URL:** `PUT /users/1`
- **Request Body:**
  ```json
  {
      "name": "Jane Doe",
      "email": "janedoe@example.com"
  }
  ```
- **Response:**
  ```json
  {
      "id": 1,
      "name": "Jane Doe",
      "email": "janedoe@example.com"
  }
  ```

### Delete User
- **Endpoint:** `DELETE /users/:id`
- **Description:** Deletes a user by their ID.
- **Example URL:** `DELETE /users/1`
- **Response:**
  ```json
  {
      "message": "User deleted successfully"
  }
  ```

## Acknowledgments
- [Go Gin](https://github.com/gin-gonic/gin) - The web framework used to build this API.

---

# Middleware
Middleware is like a helper that sits between your web server and your application. It’s a piece of code that can do things like check if a user is logged in, log requests, or modify data before the main part of your app handles it. You can think of it as a checkpoint where you can add extra rules or actions for incoming requests before they reach your main application.
