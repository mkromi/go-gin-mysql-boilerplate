# go-gin-mysql-boilerplate
> A starter project with Golang, Gin and MySQL

Golang Gin boilerplate with MySQL resource. Authenticated with Bearer and Basic authentication.

### Boilerplate structure

```
.
├── Config
│   └── Database.go
├── Controllers
│   ├── Auth.go
│   └── User.go
├── Middlewares
│   └── Auth.go
├── Models
│   ├── Auth.go
│   ├── Schema
│   │   ├── OAuthAccessToken.go
│   │   ├── OAuthClient.go
│   │   ├── OAuthRefreshToken.go
│   │   └── User.go
│   └── User.go
├── Routes
│   └── Routes.go
├── Services
│   ├── Auth.go
│   └── Response.go
├── Validations
│   ├── Auth.go
│   └── User.go
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── main.go

```

## Installation

```sh
go run main.go
```

## Usage example

`curl http://localhost:8080/users`

### Dependencies:
* [cors v1.3.1](https://github.com/gin-contrib/cors)
* [gin v1.6.3](https://github.com/gin-gonic/gin)
* [mysql v1.5.0](https://github.com/go-sql-driver/mysql)
* [gorm v1.5.0](https://github.com/jinzhu/gorm)
* [crypto v0.0.0-20191205180655-e7c4368fe9dd](https://golang.org/x/crypto)

### **License**
The **go-gin-mysql-boilerplate** is a open-source application licensed under the [MIT License](LICENSE).
