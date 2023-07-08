package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SkadisCode/meaning/config"
	"github.com/SkadisCode/meaning/delivery"
	"github.com/SkadisCode/meaning/repository"
	"github.com/SkadisCode/meaning/usecase"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Read the SQL configuration from the YAML file
	sqlConfig, err := config.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read SQL configuration: %v", err)
	}

	// Create the connection string
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		sqlConfig.Host, sqlConfig.Port, sqlConfig.Username, sqlConfig.Password, sqlConfig.Database,
	)

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")

	// Create the repository and use case instances
	dictionaryRepo := repository.NewDictionaryRepository(db)
	dictionaryUsecase := usecase.NewDictionaryUsecase(*dictionaryRepo)

	// Create the handler instance
	dictionaryHandler := delivery.NewDictionaryHandler(dictionaryUsecase)

	// Define the API routes
	router.HandleFunc("/dictionary", dictionaryHandler.SaveDictionary).Methods("POST")
	router.HandleFunc("/dictionary", dictionaryHandler.GetDictionary).Methods("GET")

	// Start the server
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
