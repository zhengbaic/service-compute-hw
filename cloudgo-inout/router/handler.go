package router

import (
	"net/http"

	"github.com/unrolled/render"
)

func rootHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}

func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string `json:"id"`
			Password string `json:"content"`
		}{Username: req.Form["username"][0], Password: req.Form["password"][0]})
	}
}
