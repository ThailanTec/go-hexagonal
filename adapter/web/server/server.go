package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ThailanTec/go-hexagonal/application/core/ports"
	"github.com/gin-gonic/gin"
)

type Webserver struct {
	Service ports.ProductServiceInterface
}

func NewWebServer() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	router := gin.Default()

	s := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
