# Devops Backend 

This repository is a Backend project using golang and and use microservice architecture.


## Project structure
```
DevOps-Backend/
├── docker-compose.yml        # Defines multi-container Docker setup (services, ports, networks, volumes)
├── README.md                 # Project documentation (how to run, build, use the system)
│
├── proto/                    # gRPC/protobuf definitions shared across services
│   ├── task-service/         
│   │   └── task.proto        # gRPC service definition for task-related operations
│   └── user-service/
│       └── user.proto        # gRPC service definition for user/auth operations
│
├── task-service/             # Microservice handling task-related business logic
│   ├── cmd/
│   │   └── main.go           # Entry point for task-service (starts the HTTP/gRPC server)
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── task-handler.go  # Contains HTTP/gRPC handler functions (e.g. CreateTask, GetTasks)
│   │   └── routes/
│   │       └── router.go     # Registers routes and applies middleware (like a mini router)
│   ├── go.mod                # Declares module name and dependencies for task-service
│   └── go.sum                # Exact versions of all dependencies used in task-service
│
├── user-service/             # Folder for all user-related services
│
│   ├── auth/                 # Authentication service (JWT, login, registration)
│   │   ├── cmd/
│   │   │   └── main.go       # Entry point for auth-service
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   │   └── auth-handler.go  # (You may want this) Logic for login, signup, token gen/verify
│   │   │   └── routes/
│   │   │       └── router.go # Registers auth routes like /login, /register
│   │   ├── go.mod            # Module definition for auth
│   │   └── go.sum
│
│   └── user/                 # User profile service (e.g. user info, update profile)
│       ├── cmd/
│       │   └── main.go       # Entry point for user-service
│       ├── internal/
│       │   ├── handlers/
│       │   │   └── user-handler.go  # Logic for user CRUD operations
│       │   └── routes/
│       │       └── router.go # Registers routes like /users, /users/:id
│       ├── go.mod            # Module for user
│       └── go.sum
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
