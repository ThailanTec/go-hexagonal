package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/ThailanTec/go-hexagonal/adapter/web/handler"
	"github.com/ThailanTec/go-hexagonal/application"
	"github.com/codegangsta/negroni"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func NewWebServer() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9001",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
