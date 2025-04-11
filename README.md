# Devops Backend 

This repository is a Backend project using golang and and use microservice architecture.


## Project structure
```
DevOps-Backend/
├── auth-service/
│   ├── cmd/
│   │   └── main.go                   # Entry point for Auth service
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── auth_handler.go       # Business logic for auth
│   │   └── routes/
│   │       └── router.go             # Sets up auth routes
│   ├── go.mod
│   └── go.sum
│
├── task-service/
│   ├── cmd/
│   │   └── main.go                   # Entry point for Task service
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── task_handler.go       # Business logic for tasks
│   │   └── routes/
│   │       └── router.go             # Sets up task routes
│   ├── go.mod
│   └── go.sum
│
├── user-service/
│   ├── cmd/
│   │   └── main.go                   # Entry point for User service
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── user_handler.go       # Business logic for user
│   │   └── routes/
│   │       └── router.go             # Sets up user routes
│   ├── go.mod
│   └── go.sum
│
├── proto/
│   ├── auth-service/
│   │   └── auth.proto                # Protobuf for auth
│   ├── task-service/
│   │   └── task.proto                # Protobuf for tasks
│   └── user-service/
│       └── user.proto                # Protobuf for users
│
├── docker-compose.yml               # Compose config to run services
└── README.md                        # Project overview
```

## Installation and run
1. Clone the repository

2. Navigate to the project directory

3. Install golang

Download link: [install golang](https://go.dev/doc/install)

4. Install dependencies:
```
go mod tidy
```

5. Running the application

Because this project using microservice architecture, if you want to run any services, you need to navigate at that service's folder and then using the below command
```
go run cmd/main.go
```
