package main

import (
	"database/sql"
	"flag"
	"github.com/sqd8642/go-gallery-hub/pkg/model" 
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8081", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:1234@localhost:5432/gogallery?sslmode=disable", "PostgreSQL DSN") //CHANGE
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Menu Singleton
	v1.HandleFunc("/images", app.createImageHandler).Methods("POST")
	v1.HandleFunc("/images/{imageId:[0-9]+}", app.getImageHandler).Methods("GET") 
	v1.HandleFunc("/images/{imageId:[0-9]+}", app.updateImageHandler).Methods("PUT") 
	v1.HandleFunc("/images/{imageId:[0-9]+}", app.deleteImageHandler).Methods("DELETE") 

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {

	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/gogallery?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
