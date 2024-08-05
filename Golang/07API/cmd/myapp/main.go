package main

import (
	"API/internal/config"
	"API/internal/database"
	"API/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//	Lectura del archivo config.yaml - configuraicones
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//	Conectarme a la DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	//	Me fijo si la tabla users existe y si no la creo
	if err := database.CreateTabla(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	//	instancio el router de gorilla mux
	router := mux.NewRouter()
	handlers.RouterHandlers(router, db)

	//	Levantamos el servidor
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
