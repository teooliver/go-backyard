package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/teooliver/go-backyard/packages/graph"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/graph/{name}", getGraph)

	http.ListenAndServe(":3000", r)
}

func getGraph(w http.ResponseWriter, r *http.Request) {
	nameParam := chi.URLParam(r, "name")
	graph, err := graph.GetGraph()
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching article %s: %v", nameParam, err)))
		return
	}

	if graph == "" {
		w.WriteHeader(404)
		w.Write([]byte("article not found"))
		return
	}

	w.Write([]byte(graph))
}
