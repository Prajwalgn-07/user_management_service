package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"user_management_service/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// Initialize configuration
	initConfig()

	// Retrieve environment variables or use default values
	dbHost := getEnvOrDefault("DATABASE_HOST", viper.GetString("database.host"))
	dbPort := getEnvOrDefault("DATABASE_PORT", viper.GetString("database.port"))
	dbUser := getEnvOrDefault("DATABASE_USER", viper.GetString("database.user"))
	dbPassword := getEnvOrDefault("DATABASE_PASSWORD", viper.GetString("database.password"))
	dbName := getEnvOrDefault("DATABASE_NAME", viper.GetString("database.dbname"))

	// Database connection setup
	db := initDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer db.Close()

	// Create user table if it doesn't exist
	const createUserTableQuery = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			email VARCHAR(100)
		);
	`
	_, err := db.Exec(createUserTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Initialize API
	apiInstance := api.NewAPI(db, router)
	apiInstance.SetupRoutes()

	// Run the server
	router.Run(":8080")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func initDB(dbHost, dbPort, dbUser, dbPassword, dbName string) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return db
}
