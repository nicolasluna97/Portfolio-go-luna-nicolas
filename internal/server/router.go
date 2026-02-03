package server

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// 1) Home handler (templates)

	mux.HandleFunc("/", HomeHandler("web/templates"))

	// 2) Static files (icons, images, etc.)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux
}
