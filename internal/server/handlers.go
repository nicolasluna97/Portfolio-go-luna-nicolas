package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title   string
	Lang    string
	IsES    bool
	IsEN    bool
	Content any
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Por ahora idioma fijo "es" (en el próximo paso lo hacemos con cookie y banderitas)
	lang := "es"

	data := PageData{
		Title: "Portfolio - Luna Nicolás",
		Lang:  lang,
		IsES:  lang == "es",
		IsEN:  lang == "en",
		Content: map[string]any{
			"Name":     "Luna Nicolás",
			"Role":     "Fullstack Developer",
			"Summary":  "Portfolio hecho en Go (SSR) + rutas internas + multi-idioma (próximo paso).",
			"Location": "Argentina",
			"Email":    "tuemail@ejemplo.com", // cambiá esto después
			"GitHub":   "https://github.com/nicolasluna97",
		},
	}

	Render(w, "home.html", data)
}

// Render renderiza una página dentro del layout base.
func Render(w http.ResponseWriter, page string, data PageData) {
	layoutPath := filepath.Join("web", "templates", "layouts", "base.html")
	pagePath := filepath.Join("web", "templates", "pages", page)

	tmpl, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(w, "Error cargando templates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Importante: base.html define {{template "base" .}}
	// y dentro incluye {{template "content" .}}
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Error renderizando template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
