package main

import (
	"cloudgo-inout/router"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", router.NewServer())
}
