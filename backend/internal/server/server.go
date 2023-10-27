package server

import (
	"database/sql"
	"log"
	"net/http"
	"zipp/internal/tools"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(db *sql.DB) {
	port := tools.EnvPortOr("3000")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleRoot)
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		handleUpload(w, r, db)
	})
	r.Get("/download", func(w http.ResponseWriter, r *http.Request) {
		handleDownload(w, r, db)
	})
	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		handleCheck(w, r, db)
	})
	r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
		handleDelete(w, r, db)
	})
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
