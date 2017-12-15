package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewServer() *negroni.Negroni {

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)
	return n
}

func initRoutes(r *mux.Router) {
	r.HandleFunc("/service/userinfo", adduser).Methods("POST")
	r.HandleFunc("/service/userinfo", getuser).Methods("GET")
	r.HandleFunc("/getallusers", getallusers).Methods("GET")
	r.HandleFunc("/", http.NotFound)
}
