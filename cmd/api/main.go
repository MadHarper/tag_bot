package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType())
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome !!!"))
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
	}
}
