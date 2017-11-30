package main

import (
	"gorm-golang/service"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", service.NewServer())
}
