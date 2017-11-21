package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	r := render.New(render.Options{
		Directory:  "views",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	m := mux.NewRouter()
	initRoutes(m, r)
	n := negroni.New()
	n.UseHandler(m)
	return n
}

func initRoutes(m *mux.Router, r *render.Render) {
	rootdir := os.Getenv("WEBROOT")
	if len(rootdir) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("can't find a root dir")
		} else {
			rootdir = root
			fmt.Println(root)
		}
	}
	m.HandleFunc("/unknown", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "this is a unknown path, plz try other path", 503)
	})
	m.HandleFunc("/login", loginHandler(r)).Methods("POST")
	m.HandleFunc("/", rootHandler(r)).Methods("GET")
	m.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(rootdir+"/public/"))))
}
