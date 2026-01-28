package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type StackItem struct {
	Key   string // clave interna (para i18n client-side si querés)
	Label string // por ahora lo mostramos tal cual
}

type ProjectCard struct {
	Title       string
	Description string
	DateRange   string
	Tech        []string
	ImageText   string // placeholder de imagen (luego lo reemplazás por imagen real)
}

type PageData struct {
	Title any
	Lang  string
	IsES  bool
	IsEN  bool

	Role string
	Name string

	Social struct {
		Github   string
		Email    string
		Linkedin string
	}

	Stack    []StackItem
	Projects []ProjectCard
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Por ahora: default ES
	lang := "es"

	var data PageData
	data.Title = "Portfolio - Luna Nicolás"
	data.Lang = lang
	data.IsES = true
	data.IsEN = false

	data.Role = "Desarrollador Full Stack"
	data.Name = "Luna Nicolás Ezequiel"

	data.Social.Github = "https://github.com/nicolasluna97"
	data.Social.Email = "mailto:nicolassluna1997@gmail.com"
	data.Social.Linkedin = "#" // lo completás después

	data.Stack = []StackItem{
		{Key: "html", Label: "HTML"},
		{Key: "css", Label: "CSS"},
		{Key: "js", Label: "JavaScript"},
		{Key: "angular", Label: "Angular"},
		{Key: "nest", Label: "NestJS"},
		{Key: "go", Label: "Go"},
		{Key: "github", Label: "GitHub"},
		{Key: "mysql", Label: "MySQL"},
		{Key: "postgres", Label: "PostgreSQL"},
		{Key: "tableplus", Label: "TablePlus"},
		{Key: "docker", Label: "Docker"},
	}

	data.Projects = []ProjectCard{
		{
			Title:       "Sistema de Facturación (Multi-tenant)",
			Description: "Plataforma de facturación e inventario con roles, productos, clientes y movimientos de stock.",
			DateRange:   "2025 — Presente",
			Tech:        []string{"Angular", "NestJS", "SQL", "Docker"},
			ImageText:   "FACTURACIÓN",
		},
		{
			Title:       "Panel de Stock & Ventas",
			Description: "Dashboard con métricas, reportes y control de inventario en tiempo real (concepto).",
			DateRange:   "2024 — 2025",
			Tech:        []string{"Go", "PostgreSQL", "HTML", "CSS"},
			ImageText:   "STOCK",
		},
		{
			Title:       "App de Gestión de Clientes",
			Description: "CRUD de clientes, proveedores y órdenes, con validaciones y filtros (concepto).",
			DateRange:   "2024",
			Tech:        []string{"Go", "MySQL", "GitHub"},
			ImageText:   "CLIENTES",
		},
	}

	Render(w, "home.html", data)
}

func Render(w http.ResponseWriter, page string, data PageData) {
	layoutPath := filepath.Join("web", "templates", "layouts", "base.html")
	pagePath := filepath.Join("web", "templates", "pages", page)

	tmpl, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(w, "Error cargando templates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Error renderizando template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
