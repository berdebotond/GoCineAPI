# Interview Test KenTech

This project serves as a backend service for a user and film management system, leveraging Go, MongoDB, and the Gin framework to provide RESTful APIs. It facilitates operations such as user authentication, signup, login, and managing film data.

## Features

- **User Authentication**: Supports user signup, login, and token-based authentication.
- **User Management**: Allows creating, retrieving, and managing users.
- **Film Management**: Enables adding, updating, retrieving, and deleting films.

## Technologies

- **Programming Language**: Go
- **Database**: MongoDB
- **Web Framework**: Gin
- **Authentication**: JWT for token management

## Getting Started

### Prerequisites

- Go (version 1.22 or later)
- MongoDB
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://yourrepositoryurl.com/project.git
   ```

2. Navigate to the project directory:

   ```bash
   cd project-directory
   ```

3. Install the Go dependencies:

   ```bash
   go mod tidy
   ```

4. Set up the environment variables:

   - `PORT`: The port on which the server will run (e.g., `8123`).
   - `MongoURI`: Your MongoDB connection string.
   - `DB_NAME`: The name of the database to use (e.g., `test`).
   - `SECRET_KEY`: A secret key used for signing JWT tokens.

   You can set these variables in your shell or use a `.env` file with [github.com/joho/godotenv](https://github.com/joho/godotenv) for loading them.

### Running the Application

Run the application using the following command:

```bash
go run ./main/main.go
The server will start, and you can interact with the APIs on http://localhost:8123.

API Endpoints

User Management
POST /users/signup: Sign up a new user.
POST /users/login: Log in a user.
Film Management
POST /films: Add a new film (requires authentication).
GET /films: List all films (requires authentication).
GET /films/:film_id: Get a specific film by ID (requires authentication).
PUT /films/:film_id: Update a specific film (requires authentication).
DELETE /films/:film_id: Delete a specific film (requires authentication).
```

### Security

This application uses JWT for handling authentication. It's highly recommended to use HTTPS in production to protect the tokens during transmission.

### Testing

This application includes two types of tests: manual API tests using Talend API Tester and automated integration tests.

#### Manual API Testing with Talend API Tester

Talend API Tester is a free standalone tool that allows you to test APIs and analyze responses. To test the application:

1. Download and install [Talend API Tester](https://www.talend.com/products/api-tester/).
2. Start the application server.
3. In Talend API Tester, import [text](talend_api_test_flow.json)
4. Send the request and analyze the response.

#### Automated Integration Testing

Automated integration tests are written in Go and can be run with the `go test` command:

1. Open a terminal.
2. Navigate to the project directory.
3. Run the following command to execute all tests in the project:

```bash
go test ./...
```