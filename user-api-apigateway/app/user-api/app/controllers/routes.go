package controllers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"user-api/app/middlewares"
)

func (server *Server) initializeRoutes() {
	server.Router.Use(middlewares.HttpMiddleware)

	//Users routes
	server.Router.HandleFunc("/users", server.CreateUser).Methods("POST")
	server.Router.HandleFunc("/users", server.GetUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
	server.Router.HandleFunc("/health", server.HealthCheck).Methods("GET")
	server.Router.HandleFunc("/getRand", server.GetRandResponseStatus).Methods("GET")
	server.Router.HandleFunc("/getUserByEmail", server.GetUserByEmail).Methods("GET")

	//Prometheus
	server.Router.Handle("/metrics", promhttp.Handler())
}
