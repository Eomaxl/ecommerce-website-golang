package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Eomaxl/ecommerce-website-golang/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", a.addr)
	return http.ListenAndServe(a.addr, router)
}
