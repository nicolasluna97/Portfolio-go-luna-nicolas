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
	URL       string
	Thumb     string
	DateRange string
	Tech      []string

	TitleES string
	TitleEN string
	DescES  string
	DescEN  string

	// Texto opcional arriba del thumb (para Creativistas u otros)
	BannerES string
	BannerEN string
}

type PageData struct {
	Title   string
	Lang    string
	Role    string
	Name    string
	HideNav bool

	Social   SocialLinks
	Stack    []StackItem
	Projects []ProjectCard
}

type ProjectPageData struct {
	Title   string
	Lang    string
	Role    string
	Name    string
	HideNav bool

	Social SocialLinks

	Slug string

	DateRange string
	Tech      []string

	TitleES string
	TitleEN string
	DescES  string
	DescEN  string

	Screenshots []string

	NotesES string
	NotesEN string
}

func HomeHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	home := filepath.Join(tplDir, "pages", "home.html")
	tpl := template.Must(template.ParseFiles(layout, home))

	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "Portfolio - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: false,
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
					// ✅ 2do proyecto
					URL:       "/projects/creativistas-web",
					Thumb:     "/static/images/projects/creativistas/creativistas-thumb.webp",
					DateRange: "2023 - 2024",
					Tech:      []string{"React", "Next.js", "Hostinger", "Express", "MongoDB Atlas", "SendGrid"},
					TitleES:   "Creativistas Web",
					TitleEN:   "Creativistas Web",
					DescES:    "Plataforma web (tests / contenidos).",
					DescEN:    "Web platform (tests / content).",

					// ✅ Dejalo vacío por ahora y después lo editás
					BannerES: "",
					BannerEN: "",
				},
				{
					// 3er proyecto por ahora sin ruta
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
		_ = tpl.ExecuteTemplate(w, "base", data)
	}
}

func InvoicingSystemHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	page := filepath.Join(tplDir, "pages", "project-invoicing.html")
	tpl := template.Must(template.ParseFiles(layout, page))

	return func(w http.ResponseWriter, r *http.Request) {
		data := ProjectPageData{
			Title:   "Sistema de Facturación - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: true, // ✅ sin navbar
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

			Screenshots: []string{
				"/static/images/projects/invoicing-system/screen-1.webp",
				"/static/images/projects/invoicing-system/screen-2.webp",
				"/static/images/projects/invoicing-system/screen-3.webp",
				"/static/images/projects/invoicing-system/screen-4.webp",
			},

			NotesES: "",
			NotesEN: "",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = tpl.ExecuteTemplate(w, "base", data)
	}
}

func CreativistasWebHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	page := filepath.Join(tplDir, "pages", "project-creativistas.html")
	tpl := template.Must(template.ParseFiles(layout, page))

	return func(w http.ResponseWriter, r *http.Request) {
		data := ProjectPageData{
			Title:   "Creativistas Web - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: true, // ✅ sin navbar
			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "#",
			},

			Slug:      "creativistas-web",
			DateRange: "2023 - 2024",
			Tech:      []string{"React", "Next.js"},

			TitleES: "Creativistas Web",
			TitleEN: "Creativistas Web",
			DescES:  "Fui contactado para resolver una serie de problemas técnicos en una aplicación web utilizada para la realización de tests psicológicos en un estudio privado. Por motivos de privacidad, en este proyecto se muestra únicamente una captura representativa del sistema.",
			DescEN:  "I was contacted to resolve several technical issues with a website that administered psychological tests for a private study. Due to website privacy concerns, I will only include one screenshot.",

			Screenshots: []string{
				"/static/images/projects/creativistas/creativistas-2.webp",
			},

			NotesES: "",
			NotesEN: "",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = tpl.ExecuteTemplate(w, "base", data)
	}
}
