package main

import (
	"net/http"

	"github.com/ngvanthanggit/GolangHttpServer.git/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8000", srv)
}
