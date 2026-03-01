package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

/* ==========
   MODELOS
   ========== */

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
}

type PageData struct {
	Title string
	Lang  string
	Role  string
	Name  string

	Social SocialLinks

	StackPrimary []StackItem
	StackOther   []StackItem

	Projects []ProjectCard
	HideNav  bool
}

/* Sección individual del proyecto (para proyecto 1) */
type ProjectPoint struct {
	ES string
	EN string
}

type ProjectSection struct {
	TitleES string
	TitleEN string
	Image   string
	AltES   string
	AltEN   string
	Points  []ProjectPoint
}

type ProjectPageData struct {
	Title string
	Lang  string
	Role  string
	Name  string

	Social  SocialLinks
	HideNav bool

	Slug      string
	HeroImage string
	DateRange string
	Tech      []string

	TitleES string
	TitleEN string
	DescES  string
	DescEN  string

	Screenshots []string
	Sections    []ProjectSection
}

/* ==========
   HOME
   ========== */

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
			HideNav: false, // ✅ Home SI muestra navbar
			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "https://www.linkedin.com/",
			},
			StackPrimary: []StackItem{
				{Key: "angular", Label: "Angular"},
				{Key: "nestjs", Label: "NestJS"},
				{Key: "postgresql", Label: "PostgreSQL"},
				{Key: "docker", Label: "Docker"},
				{Key: "github", Label: "GitHub"},
			},
			StackOther: []StackItem{
				{Key: "go", Label: "Go"},
				{Key: "javascript", Label: "JavaScript"},
				{Key: "mysql", Label: "MySQL"},
			},

			Projects: []ProjectCard{
				{
					URL:       "/projects/invoicing-system",
					Thumb:     "/static/images/projects/invoicing-system/thumb.webp",
					DateRange: "2025 - 2026",
					Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},
					TitleES:   "Sistema de Facturación (Multi-tenant)",
					TitleEN:   "Invoicing System (Multi-tenant)",
					DescES:    "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
					DescEN:    "Multi-tenant invoicing and inventory platform with roles, customers, products, and stock movement tracking.",
				},
				{
					URL:       "/projects/creativistas-web",
					Thumb:     "/static/images/projects/creativistas/creativistas-thumb.webp",
					DateRange: "2025",
					Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},
					TitleES:   "Creativistas Web",
					TitleEN:   "Creativistas Web",
					DescES:    "Web de tests psicológicos, con envío de resultados por email.",
					DescEN:    "Psychological tests, with automated email delivery of results.",
				},
				{
					URL:       "/projects/tienda-nube",
					TitleES:   "Tienda Online (TiendaNube)",
					TitleEN:   "Online Store (TiendaNube)",
					DescES:    "E-commerce en TiendaNube: personalización del theme, configuración de catálogo/variantes y optimización del flujo de compra para autogestión.",
					DescEN:    "TiendaNube e-commerce: theme customization, catalog/variants setup, and checkout flow optimization for client self-management.",
					Thumb:     "/static/images/projects/tiendanube/thumb-tiendanube.webp",
					DateRange: "2025 — 2026",
					Tech:      []string{"TiendaNube", "UI/UX", "CSS", "HTML"},
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

/* ==========
   PROYECTO 1: INVOICING
   ========== */

func InvoicingSystemHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	page := filepath.Join(tplDir, "pages", "project-invoicing.html")
	tpl := template.Must(template.ParseFiles(layout, page))

	return func(w http.ResponseWriter, r *http.Request) {
		data := ProjectPageData{
			Title:   "Invoicing System - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: true,

			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "https://www.linkedin.com/",
			},

			Slug:      "invoicing-system",
			HeroImage: "/static/images/projects/invoicing-system/thumb.webp",
			DateRange: "2025 - 2026",
			Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},

			TitleES: "Sistema de Facturación (Multi-tenant)",
			TitleEN: "Invoicing System (Multi-tenant)",
			DescES:  "Plataforma de facturación e inventario con roles, clientes, productos, movimientos de stock y estadísticas para empresas o autonomos. La página web posse actulaización de stock en tiempo real además modo oscuro y claro. Actualmente en desarrollo en beta 2.0, cuando se sumen las funcionalidades faltantes, se subirá a producción.",
			DescEN:  "Multi-tenant invoicing and inventory platform with roles, customers, products, stock movements, and analytics for businesses and self-employed professionals. It provides real-time stock updates and supports both dark and light themes. Currently in beta 2.0; once the remaining features are completed, it will be released to production.",

			Screenshots: []string{
				"/static/images/projects/invoicing-system/inventory-screen.webp",
				"/static/images/projects/invoicing-system/sales-screen.webp",
				"/static/images/projects/invoicing-system/statistics-screen.webp",
				"/static/images/projects/invoicing-system/clients-screen.webp",
			},

			Sections: []ProjectSection{
				{
					TitleES: "Inventario",
					TitleEN: "Inventory",
					Image:   "/static/images/projects/invoicing-system/inventory-screen.webp",
					AltES:   "Gestión de inventario",
					AltEN:   "Inventory management",
					Points: []ProjectPoint{
						{
							ES: "Permite crear productos mediante formularios modales, asignando categorías, stock inicial, precio de compra y múltiples precios de venta por producto.",
							EN: "Allows creating products via modal forms, assigning categories, initial stock, purchase cost, and multiple selling prices per product.",
						},
						{
							ES: "Habilita edición masiva de productos seleccionados, con validaciones por campo y feedback visual de errores.",
							EN: "Enables bulk editing of selected products, with per-field validations and visual error feedback.",
						},
						{
							ES: "Soporta eliminación segura de uno o varios productos mediante modales de confirmación para evitar acciones accidentales.",
							EN: "Supports safe deletion of one or multiple products through confirmation modals to prevent accidental actions.",
						},
					},
				},
				{
					TitleES: "Ventas",
					TitleEN: "Sales",
					Image:   "/static/images/projects/invoicing-system/sales-screen.webp",
					AltES:   "Flujo de ventas",
					AltEN:   "Sales flow",
					Points: []ProjectPoint{
						{
							ES: "Ofrece un flujo de venta rápido, pensado para uso operativo, donde se seleccionan productos, cantidades y lista de precios.",
							EN: "Provides a fast, operations-oriented sales flow where products, quantities, and a price list are selected.",
						},
						{
							ES: "Soporta múltiples precios por producto, permitiendo calcular totales según el precio aplicado en cada venta.",
							EN: "Supports multiple prices per product, calculating totals based on the price applied to each sale.",
						},
						{
							ES: "Al confirmar la operación, descuenta automáticamente el stock, impactando en inventario en tiempo real.",
							EN: "On confirmation, it automatically decreases stock, updating inventory in real time.",
						},
					},
				},
				{
					TitleES: "Estadísticas",
					TitleEN: "Analytics",
					Image:   "/static/images/projects/invoicing-system/statistics-screen.webp",
					AltES:   "Panel de estadísticas",
					AltEN:   "Analytics dashboard",
					Points: []ProjectPoint{
						{
							ES: "Permite analizar resultados por día, semana, mes o año, mediante filtros de fecha dinámicos.",
							EN: "Allows analyzing results by day, week, month, or year using dynamic date filters.",
						},
						{
							ES: "Muestra KPIs principales como ventas totales, productos vendidos y ganancias del período seleccionado.",
							EN: "Shows key KPIs such as total sales, items sold, and profit for the selected period.",
						},
						{
							ES: "Integra gráficos para visualizar la evolución de la recaudación y facilitar la toma de decisiones.",
							EN: "Includes charts to visualize revenue trends and support decision-making.",
						},
					},
				},
				{
					TitleES: "Clientes",
					TitleEN: "Customers",
					Image:   "/static/images/projects/invoicing-system/clients-screen.webp",
					AltES:   "Gestión de clientes",
					AltEN:   "Customer management",
					Points: []ProjectPoint{
						{
							ES: "Permite registrar y administrar clientes, centralizando la información de contacto y relación comercial.",
							EN: "Lets you register and manage customers, centralizing contact details and business relationship data.",
						},
						{
							ES: "Facilita la asociación de ventas a clientes, mejorando el seguimiento histórico de operaciones.",
							EN: "Enables linking sales to customers, improving historical tracking of operations.",
						},
						{
							ES: "Ofrece un listado estructurado para consulta y gestión rápida desde el sistema interno.",
							EN: "Provides a structured list for quick lookup and management inside the system.",
						},
					},
				},

				// Para las secciones sin imagen: igual, pero Image/Alt pueden quedar vacíos.
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/* ==========
   PROYECTO 2: CREATIVISTAS
   ========== */

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
			HideNav: true,

			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "https://www.linkedin.com/",
			},

			Slug:      "creativistas",
			HeroImage: "/static/images/projects/creativistas/creativistas-thumb.webp",
			DateRange: "2025",
			Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},

			TitleES: "Creativistas Web",
			TitleEN: "Creativistas Web",
			DescES:  "Web de tests psicológicos, con envío de resultados por email.",
			DescEN:  "Psychological tests, with email delivery of results.",

			Screenshots: []string{
				"/static/images/projects/creativistas/creativistas-2.webp",
			},
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/* ==========
PROYECTO 3: TIENDANUBE
========== */

func TiendaNubeHandler(tplDir string) http.HandlerFunc {
	layout := filepath.Join(tplDir, "layouts", "base.html")
	page := filepath.Join(tplDir, "pages", "project-tiendanube.html")
	tpl := template.Must(template.ParseFiles(layout, page))

	return func(w http.ResponseWriter, r *http.Request) {
		data := ProjectPageData{
			Title:   "TiendaNube - Luna Nicolás",
			Lang:    "es",
			Role:    "Desarrollador Full Stack",
			Name:    "Luna Nicolás Ezequiel",
			HideNav: true,

			Social: SocialLinks{
				Github:   "https://github.com/nicolasluna97",
				Email:    "mailto:nicolassluna1997@gmail.com",
				Linkedin: "https://www.linkedin.com/",
			},

			Slug:      "tienda-nube",
			HeroImage: "/static/images/projects/tiendanube/tiendanube-website.webp",
			DateRange: "2025 - 2026",
			Tech:      []string{"TiendaNube", "UI/UX", "HTML", "CSS"},

			TitleES: "Tienda Online (TiendaNube)",
			TitleEN: "Online Store (TiendaNube)",
			DescES:  "Implementación y personalización de tienda en TiendaNube.",
			DescEN:  "Implementation and customization of a TiendaNube store.",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
