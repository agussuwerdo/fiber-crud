# Go Fiber + MongoDB CRUD Application

This project is a basic CRUD application built with [Go Fiber](https://gofiber.io/), [MongoDB](https://www.mongodb.com/), and JWT-based authentication. It allows users to manage items (create, read, update, delete) in a MongoDB database and includes JWT-based user authentication.

## Features

- JWT Authentication for secured routes
- CRUD operations for items stored in MongoDB
- Environment-based configuration using `.env`
- Fiber web framework for fast HTTP handling
- MongoDB as the database

## Prerequisites

Before running this project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16+)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Git](https://git-scm.com/)
- [Go Fiber](https://gofiber.io/)

## Project Structure

```bash
fiber-crud/
│
├── .env                    # Environment configuration file
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies lock file
├── main.go                  # Main entry point for the Go Fiber application
├── handlers/                # Contains logic for handling routes
│   ├── item_handler.go       # Handlers for CRUD operations on items
│   ├── login_handler.go      # Handlers for authentication
│   └── mongo_handler.go      # MongoDB initialization logic
├── models/                  # Contains the models for the app
│   └── item.go               # Item model definition
└── README.md                # This readme file
