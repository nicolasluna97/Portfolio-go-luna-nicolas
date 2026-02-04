package server

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Home
	mux.HandleFunc("/", HomeHandler("web/templates"))

	// Proyecto específico
	mux.HandleFunc("/projects/invoicing-system", InvoicingSystemHandler("web/templates"))

	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux
}
