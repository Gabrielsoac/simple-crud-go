package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
);

func SetupDataBase() *sql.DB {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file");
	}

	dbHost := os.Getenv("DB_HOST");
	dbPort := os.Getenv("DB_PORT");
	dbName := os.Getenv("DB_NAME");
	dbUser := os.Getenv("DB_USERNAME");
	dbPassword := os.Getenv("DB_PASSWORD");
	
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName);

	fmt.Println(connectionStr)
	dbConnection, error := sql.Open("postgres", connectionStr);
	if error != nil {
		log.Fatal("Error connecting db")
		os.Exit(1)
	}

	error = dbConnection.Ping();
	if error != nil {
		log.Fatal("Db connection not initialized");
		os.Exit(1)
	}

	fmt.Println("Success conneting db");
	return dbConnection;
}