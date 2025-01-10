# MongoDB Golang REST API

A RESTful API built using Golang and MongoDB to manage employee data. This project includes endpoints for creating, reading, updating, and deleting employees, and unit tests for database operations.

You can follow this wonderful youtube playlist [Youtube Playlist](https://youtube.com/playlist?list=PL-92QVNKZDMyfu1VkmXnc1STRi0MpUJ4-&si=1Mx_Zyz-uyK0V-nB)


---

## Features

- CRUD operations for employee data
- MongoDB integration using `mongo-driver`
- Environment variable support using `.env`
- Middleware support with `chi` router
- Modular and testable codebase
- Unit tests for database operations

---

## Requirements

- Go 1.23 or higher
- MongoDB instance
- `.env` file for configuration

---

## Project Structure

```plaintext
MONGODB_GOLANG_REST_API
├── model
│   └── employee.go           # Employee data model
├── repository
│   ├── employee.go           # Repository for MongoDB operations
│   └── employee_test.go      # Unit tests for MongoDB repository
├── usecase
│   └── employee.go           # Business logic and use case implementation
├── .env                      # Environment variables configuration
├── go.mod                    # Go module dependencies
├── go.sum                    # Dependency checksums
├── main.go                   # Entry point of the application
└── README.md                 # Project documentation
```

---

## Environment Variables

Create a `.env` file in the root directory with the following content:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=companydb
COLLECTION_NAME=employee
```

Modify these values to suit your MongoDB setup.

---

## Installation and Setup

### 1. Clone the repository

```bash
git clone https://github.com/naheedrayan/mongodb_golang_rest_api.git
cd mongodb_golang_rest_api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

Start the server using:

```bash
go run main.go
```

The server will start on `http://localhost:3000`.

---

## API Endpoints

### Base URL: `http://localhost:3000`

| Method | Endpoint                 | Description                      |
|--------|--------------------------|----------------------------------|
| POST   | `/employees`             | Create a new employee           |
| GET    | `/employees`             | Get all employees               |
| GET    | `/employees/{id}`        | Get an employee by ID           |
| PUT    | `/employees/{id}`        | Update an employee by ID        |
| DELETE | `/employees/{id}`        | Delete an employee by ID        |
| DELETE | `/employees`             | Delete all employees            |
| GET    | `/`                      | Welcome message                 |

---

## Code Overview

### `main.go`

The entry point of the application. Initializes the MongoDB connection, sets up the `chi` router, and defines the API endpoints.

### `model/employee.go`

Defines the `Employee` struct used as the data model for MongoDB operations.

### `repository/employee.go`

Contains MongoDB repository logic, such as creating, retrieving, updating, and deleting employee records.

### `repository/employee_test.go`

Unit tests for the MongoDB repository functions to ensure correct functionality.

### `usecase/employee.go`

Implements the business logic for employee-related operations, delegating database interactions to the repository layer.

---

## Unit Tests

Unit tests are located in `repository/employee_test.go`.

Run the tests using:

```bash
go test ./repository
```

---

## Dependencies

- [MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
- [Chi Router](https://github.com/go-chi/chi)
- [Godotenv](https://github.com/joho/godotenv)

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
