package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type SocialLinks struct {
	Github   string
	Email    string
	Linkedin string
}

type StackItem struct {
	Key   string
	Label string
}

/*
✅ Ahora cada proyecto tiene:
- TitleES / TitleEN
- DescES / DescEN
Así el default ES se ve bien, y EN aparece cuando el usuario cambia idioma.
*/
type ProjectCard struct {
	Thumb     string
	DateRange string
	Tech      []string

	TitleES string
	TitleEN string
	DescES  string
	DescEN  string
}

type PageData struct {
	Title string
	Lang  string
	Role  string
	Name  string

	Social SocialLinks
	Stack  []StackItem

	Projects []ProjectCard
}

func HomeHandler(tplDir string) http.HandlerFunc {
	// Parse templates
	layout := filepath.Join(tplDir, "layouts", "base.html")
	home := filepath.Join(tplDir, "pages", "home.html")

	tpl := template.Must(template.ParseFiles(layout, home))

	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Portfolio - Luna Nicolás",
			Lang:  "es",
			Role:  "Desarrollador Full Stack",
			Name:  "Luna Nicolás Ezequiel",
			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "#",
			},
			Stack: []StackItem{
				{Key: "html", Label: "HTML"},
				{Key: "css", Label: "CSS"},
				{Key: "javascript", Label: "JavaScript"},
				{Key: "angular", Label: "Angular"},
				{Key: "nestjs", Label: "NestJS"},
				{Key: "go", Label: "Go"},
				{Key: "github", Label: "GitHub"},
				{Key: "mysql", Label: "MySQL"},
				{Key: "postgresql", Label: "PostgreSQL"},
				{Key: "docker", Label: "Docker"},
			},
			Projects: []ProjectCard{
				{
					Thumb:     "/static/images/projects/invoicing-system/thumb.webp",
					DateRange: "2024 - 2025",
					Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},

					TitleES: "Sistema de Facturación (Multi-tenant)",
					TitleEN: "Invoicing System (Multi-tenant)",
					DescES:  "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
					DescEN:  "Multi-tenant invoicing and inventory management system with roles, customers, products and stock movements.",
				},
				{
					Thumb:     "/static/images/projects/stock-dashboard/thumb.webp",
					DateRange: "2024",
					Tech:      []string{"Angular", "NestJS", "Socket.IO"},

					TitleES: "Panel de Stock & Ventas",
					TitleEN: "Stock & Sales Dashboard",
					DescES:  "Dashboard con métricas, reportes y control de inventario en tiempo real (concepto).",
					DescEN:  "Dashboard with metrics, reports and real-time inventory control (concept).",
				},
				{
					Thumb:     "/static/images/projects/clients-app/thumb.webp",
					DateRange: "2023 - 2024",
					Tech:      []string{"Angular", "NestJS", "MySQL"},

					TitleES: "App de Gestión de Clientes",
					TitleEN: "Client Management App",
					DescES:  "CRUD de clientes, proveedores y órdenes, con validaciones y filtros (concepto).",
					DescEN:  "CRUD for clients, suppliers and orders, with validations and filters (concept).",
				},
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = tpl.ExecuteTemplate(w, "base", data)
	}
}
