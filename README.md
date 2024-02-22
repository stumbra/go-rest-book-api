# go-rest-book-api

This repository serves as a learning project for creating a REST API using the `gin-gonic/gin` package in Go.

## Description

The primary purpose of this project is to demonstrate CRUD operations using a RESTful architecture for managing books. It showcases how to create endpoints for creating, reading, updating, and deleting books.

## Features

- CRUD operations for managing books
- Integration of Swagger documentation generation for API endpoints
- Hot reload feature with Air package for seamless development experience
- Structured project layout with separate folders for controllers, mocks, and models

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/stumbra/go-rest-book-api.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-rest-book-api
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Start the server with hot reload:

   ```bash
   air
   ```

2. Access the API documentation:

   Navigate to [http://localhost:9090/swagger/index.html](http://localhost:9090/swagger/index.html) in your web browser.

## Project Structure

The project follows a structured layout for better organization:

```
.
├── controllers    # Controllers for handling HTTP requests
├── mocks          # Mocked data
├── models         # Data models for the application
├── README.md      # This README file
└── main.go        # Entry point of the application
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests for any improvements or features you'd like to see in this project.

## License

This project is licensed under the [MIT License](LICENSE).
