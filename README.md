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
Go-RESTful-API-Exercise/
├── controllers
├── database
├── docs
├── go.mod
├── go.sum
└── main.go
```

- `controllers/`: Contains the controllers that handle incoming requests.
- `database/`: Contains the database configuration and migrations.
- `docs/`: Contains the Swagger API documentation.
- `go.mod`: The Go module file that lists the dependencies of this module.
- `go.sum`: The Go module checksum file that lists the cryptographic checksums of the dependencies.
- `main.go`: The entry point of the application.

## Usage

To use this module, you need to have Go installed. Clone this repository and navigate to the root folder of the module. Run the following command to start the server:

```
go run main.go
```

The server will start on port 8080. You can access the Swagger API documentation by navigating to `http://localhost:8080/swagger/index.html` in your browser.