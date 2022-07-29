package server

import (
	"fmt"
	"golang-app/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router   *mux.Router
	postgres *db.MyDatabase
}

func NewServer() {
	p := db.OpenConnection()
	s := &server{router: mux.NewRouter(),
		postgres: p}
	s.routes()
	fmt.Println("server is running on: 3000")

	log.Fatalln(http.ListenAndServe(":3000", s.router))
}
