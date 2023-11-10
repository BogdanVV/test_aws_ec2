package server

import (
	"fmt"
	"net/http"
	"time"
)

var port = 9999

type Server struct {
	Port int
}

func NewServer() *http.Server {
	newServer := &Server{
		Port: port,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.Port),
		IdleTimeout:  time.Minute,
		Handler:      newServer.InitRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
