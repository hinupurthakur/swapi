package routes

import (
	"net/http"

	"github.com/hinupurthakur/swapi/api"

	"github.com/gorilla/mux"
)

func CreateRoutes() http.Handler {
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()
	r.Handle("/", Routes(r))
	// wrap all routes
	return r
}

func Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", api.HealthCheck).Methods("GET")
	return r
}
