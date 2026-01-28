package server

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Rutas principales
	mux.HandleFunc("/", Home)

	// (Más adelante) proyectos:
	// mux.HandleFunc("/projects", Projects)
	// mux.HandleFunc("/projects/", ProjectDetail)

	// Archivos estáticos (si los agregamos después):
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux
}
