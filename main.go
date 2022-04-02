package main

import (
	"go-chi/router"
	"net/http"
)

func main() {
	r := router.Routers()
	http.ListenAndServe(":8080", r)
}
