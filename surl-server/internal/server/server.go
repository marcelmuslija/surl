package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"surl-server/internal/config"
	"surl-server/internal/db"
	"surl-server/internal/server/short"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
	Config *config.Config
}

func New() (*Server, error) {
	c, err := config.Get()
	if err != nil {
		return nil, err
	}

	db, err := db.Connect(c)
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	shortUrlHandler := short.NewHandler(db)
	r.HandleFunc("/api/short", shortUrlHandler.Shorten).Methods("POST")
	r.HandleFunc("/api/short/{token}", shortUrlHandler.Redirect).Methods("GET")
	r.HandleFunc("/api/short/{token}", shortUrlHandler.Shorten).Methods("DELETE")

	return &Server{
		Router: r,
		DB:     db,
		Config: c,
	}, nil
}
