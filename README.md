# Go-RESTful-API-Exercise

This is a Go module for a RESTful API exercise. It is written in Go version 1.20 and uses the following dependencies:

- github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
- github.com/gin-gonic/gin v1.9.0
- github.com/golang-jwt/jwt v3.2.2+incompatible
- gorm.io/driver/postgres v1.5.0
- gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11

## Folder Structure

The folder structure of this module is as follows:

```
Go RESTful-API Exercise/
├── controllers
├── database
├── docs
├── middlewares
├── models
├── router
├── go.mod
├── go.sum
├── helpers
├── main.go
├── README.md
└── req.http
```

- `controllers/`: Contains the controllers that handle incoming requests.
- `database/`: Contains the database configuration and migrations.
- `docs/`: Contains the Swagger API documentation.
- `middlewares/`: Contains the middleware functions that are used in the application. Middleware functions are functions that are invoked before or after the main handler function of an API endpoint. They are used for tasks such as authentication, logging, and error handling.
- `models/`: Contains the data models that represent the entities in the application's domain. These models typically map to tables in a database and define the structure and relationships of the data.
- `router/`: Contains the definition of the API routes and their corresponding handler functions. The router is responsible for matching incoming requests to the appropriate handler function.
- `go.mod`: The Go module file that lists the dependencies of this module.
- `go.sum`: The Go module checksum file that lists the cryptographic checksums of the dependencies.
- `main.go`: The entry point of the application.

## Usage

To use this module, you need to have Go installed. Clone this repository and navigate to the root folder of the module. Run the following command to start the server:

```
go run main.go
```

The server will start on port 8080. You can access the Swagger API documentation by navigating to `http://localhost:8080/swagger/index.html` in your browser.