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

type ProjectCard struct {
	URL         string
	Thumb       string
	DateRange   string
	Tech        []string
	TitleES     string
	TitleEN     string
	DescES      string
	DescEN      string
	Screenshots []string // ✅ IMPORTANTE (esto te faltaba)
}

type PageData struct {
	Title   string
	Lang    string
	Role    string
	Name    string
	HideNav bool // ✅ para ocultar navbar si querés

	Social SocialLinks
	Stack  []StackItem

	Projects []ProjectCard
}

type ProjectPageData struct {
	Title   string
	Lang    string
	Role    string
	Name    string
	HideNav bool // ✅ ocultar navbar en rutas /projects/*

	Social SocialLinks

	Slug        string
	DateRange   string
	Tech        []string
	TitleES     string
	TitleEN     string
	DescES      string
	DescEN      string
	Screenshots []string
}

func HomeHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	home := filepath.Join(tplDir, "pages", "home.html")

	tpl, err := template.ParseFiles(layout, home)
	if err != nil {
		panic(err)
	}

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
				{Key: "html5", Label: "HTML"},
				{Key: "css3", Label: "CSS"},
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
					URL:       "/projects/invoicing-system",
					Thumb:     "/static/images/projects/invoicing-system/thumb.webp",
					DateRange: "2024 - 2025",
					Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},
					TitleES:   "Sistema de Facturación (Multi-tenant)",
					TitleEN:   "Invoicing System (Multi-tenant)",
					DescES:    "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
					DescEN:    "Multi-tenant invoicing and inventory management system with roles, customers, products and stock movements.",
				},
				{
					URL:       "#",
					Thumb:     "/static/images/projects/stock-dashboard/thumb.webp",
					DateRange: "2024",
					Tech:      []string{"Angular", "NestJS", "Socket.IO"},
					TitleES:   "Panel de Stock & Ventas",
					TitleEN:   "Stock & Sales Dashboard",
					DescES:    "Dashboard con métricas, reportes y control de inventario en tiempo real (concepto).",
					DescEN:    "Dashboard with metrics, reports and real-time inventory control (concept).",
				},
				{
					URL:       "#",
					Thumb:     "/static/images/projects/clients-app/thumb.webp",
					DateRange: "2023 - 2024",
					Tech:      []string{"Angular", "NestJS", "MySQL"},
					TitleES:   "App de Gestión de Clientes",
					TitleEN:   "Client Management App",
					DescES:    "CRUD de clientes, proveedores y órdenes, con validaciones y filtros (concepto).",
					DescEN:    "CRUD for clients, suppliers and orders, with validations and filters (concept).",
				},
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func InvoicingSystemHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	page := filepath.Join(tplDir, "pages", "project-invoicing.html")

	tpl, err := template.ParseFiles(layout, page)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		data := ProjectPageData{
			Title:   "Sistema de Facturación - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: true, // ✅ ACA se oculta el navbar SOLO en esta ruta

			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "#",
			},

			Slug:      "invoicing-system",
			DateRange: "2024 - 2025",
			Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},

			TitleES: "Sistema de Facturación (Multi-tenant)",
			TitleEN: "Invoicing System (Multi-tenant)",
			DescES:  "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
			DescEN:  "Multi-tenant invoicing and inventory management system with roles, customers, products and stock movements.",

			// ✅ TUS RUTAS REALES (screen-1.webp etc)
			Screenshots: []string{
				"/static/images/projects/invoicing-system/screen-1.webp",
				"/static/images/projects/invoicing-system/screen-2.webp",
				"/static/images/projects/invoicing-system/screen-3.webp",
				"/static/images/projects/invoicing-system/screen-4.webp",
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
