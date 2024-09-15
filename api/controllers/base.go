package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	stein "github.com/nasrul21/go-stein"
)

type Server struct {
	STHQ   *stein.Stein
	Router *mux.Router
}

func (server *Server) Initialize(SteinUrl, SteinUser, SteinPassword string) {
	server.STHQ = stein.NewClient(
		SteinUrl,
		&stein.Option{
			Username: SteinUser,
			Password: SteinPassword,
			Timeout:  20 * time.Second,
		})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
