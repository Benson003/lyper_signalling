package server

import (
	"net/http"
	"time"
)
type Server struct {}

func NewServer() *http.Server {
	server := http.Server{
		Addr:         "0.0.0.0:4000",
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 15,
	}
	return &server
}
