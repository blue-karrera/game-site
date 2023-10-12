package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

    "game-site/internal"
)

func root(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome"))
}

func network_test() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
    r.Get("/", root)
	http.ListenAndServe(":3000", r)
}

func bridge_test() {
}

func main() {
    internal.TestHS()
}


