# Interview Test KenTech

Welcome to GoCineAPI, a comprehensive API designed for managing films and users lists, built with Go. This project was specifically created as part of a backend developer interview test for KenTech, showcasing best practices in API development using Go.

GoCineAPI leverages JWT (JSON Web Token) for secure authentication, ensuring that only registered users can access and interact with the movie-related endpoints. Whether it's registering a new user, logging in, creating unique films, or managing a personal favorites list, GoCineAPI covers all bases, providing a solid foundation for any developer looking to implement or expand their understanding of such systems.

For KenTech Interview
This project was developed with a specific purpose—to address the requirements of a backend developer position at KenTech, showcasing a practical application of Go in solving real-world problems. It reflects a deep understanding of API development, secure authentication practices, and the efficient management of film data.

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
   git clone https://github.com/berdebotond/GoCineAPI.git
   ```

2. Navigate to the project directory:

   ```bash
   cd GoCineAPI
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
3. In Talend API Tester, import talend_api_test_flow.json
4. Send the request and analyze the response.

#### Automated Integration Testing

Automated integration tests are written in Go and can be run with the `go test` command:

1. Open a terminal.
2. Navigate to the project directory.
3. Run the following command to execute all tests in the project:

```bash
docker-compose -f mongodb.yaml up -d
go test ./...
```
