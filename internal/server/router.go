package server

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)

	// Más adelante:
	// mux.HandleFunc("/projects", Projects)
	// mux.HandleFunc("/projects/", ProjectDetail)

	return mux
}
