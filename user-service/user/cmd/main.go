package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/minhphuc2544/DevOps-Backend/user-service/user/internal/routes"
)



func main() {
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to get current working directory: %v", err)
    }

    // Construct the absolute path to the .env file
    envPath := filepath.Join(cwd,"..","..", ".env") 
    log.Printf("Attempting to load .env file from: %s", envPath)

    // Load the .env file
    err = godotenv.Load(envPath)
    if err != nil {
        log.Fatalf("Error loading .env file from %s: %v", envPath, err)
    }
    dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DATABASE")

    // Establish a database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    // Verify the connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }
    log.Println("Successfully connected to the database.")

    router := routes.SetupRoutes(db) // Setup the routes
    log.Println("Starting server on :8080...")
    // Start the server on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}