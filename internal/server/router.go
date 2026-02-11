package server

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Home
	mux.HandleFunc("/", HomeHandler("web/templates"))

	// Proyectos (rutas puntuales)
	mux.HandleFunc("/projects/invoicing-system", InvoicingSystemHandler("web/templates"))
	mux.HandleFunc("/projects/creativistas-web", CreativistasWebHandler("web/templates"))
	mux.HandleFunc("/projects/tienda-nube", TiendaNubeHandler("web/templates"))

	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux
}
