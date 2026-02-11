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
	Stack  []StackItem

	Projects []ProjectCard

	// ✅ FIX CLAVE: base.html usa .HideNav
	HideNav bool
}

/* Sección individual del proyecto (para proyecto 1) */
type ProjectSection struct {
	Title  string
	Image  string
	Alt    string
	Points []string
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
					DateRange: "2025 - actualidad",
					Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},
					TitleES:   "Sistema de Facturación (Multi-tenant)",
					TitleEN:   "Invoicing System (Multi-tenant)",
					DescES:    "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
					DescEN:    "Multi-tenant invoicing and inventory platform with roles, customers, products and stock movements.",
				},
				{
					URL:       "/projects/creativistas-web",
					Thumb:     "/static/images/projects/creativistas/creativistas-thumb.webp",
					DateRange: "2025 - 2026",
					Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},
					TitleES:   "Creativistas Web",
					TitleEN:   "Creativistas Web",
					DescES:    "Web de tests psicológicos (Big 5) con envío de resultados por email.",
					DescEN:    "Psychological tests (Big 5) with email delivery of results.",
				},
				{
					URL:       "/projects/tienda-nube",
					TitleES:   "Tienda Online (TiendaNube)",
					TitleEN:   "Online Store (TiendaNube)",
					DescES:    "E-commerce en TiendaNube: personalización del theme, configuración de catálogo/variantes y optimización del flujo de compra para autogestión.",
					DescEN:    "TiendaNube e-commerce: theme customization, catalog/variants setup, and checkout flow optimization for client self-management.",
					Thumb:     "/static/images/projects/tiendanube/thumb-tiendanube.webp",
					DateRange: "2026 — Actualidad",
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
			DateRange: "2024 - 2025",
			Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},

			TitleES: "Sistema de Facturación (Multi-tenant)",
			TitleEN: "Invoicing System (Multi-tenant)",
			DescES:  "Plataforma de facturación e inventario con roles, clientes, productos, movimientos de stock y estadísticas para empresas o autonomos. La página web posse actulaización de stock en tiempo real además modo oscuro y claro. Actualmente en desarrollo en beta 2.0, cuando se sumen las funcionalidades faltantes, se subirá a producción.",
			DescEN:  "Multi-tenant invoicing and inventory platform with roles, customers, products and stock movements.",

			Screenshots: []string{
				"/static/images/projects/invoicing-system/inventory-screen.webp",
				"/static/images/projects/invoicing-system/sales-screen.webp",
				"/static/images/projects/invoicing-system/statistics-screen.webp",
				"/static/images/projects/invoicing-system/clients-screen.webp",
			},

			Sections: []ProjectSection{
				{
					Title: "Inventario",
					Image: "/static/images/projects/invoicing-system/inventory-screen.webp",
					Alt:   "Gestión de inventario",
					Points: []string{
						"Permite crear productos mediante formularios modales, asignando categorías, stock inicial, precio de compra y múltiples precios de venta por producto.",
						"Habilita edición masiva de productos seleccionados, con validaciones por campo y feedback visual de errores.",
						"Soporta eliminación segura de uno o varios productos mediante modales de confirmación para evitar acciones accidentales.",
					},
				},
				{
					Title: "Ventas",
					Image: "/static/images/projects/invoicing-system/sales-screen.webp",
					Alt:   "Flujo de ventas",
					Points: []string{
						"Ofrece un flujo de venta rápido, pensado para uso operativo, donde se seleccionan productos, cantidades y lista de precios.",
						"Soporta múltiples precios por producto, permitiendo calcular totales según el precio aplicado en cada venta.",
						"Al confirmar la operación, descuenta automáticamente el stock, impactando en inventario en tiempo real.",
					},
				},
				{
					Title: "Estadísticas",
					Image: "/static/images/projects/invoicing-system/statistics-screen.webp",
					Alt:   "Panel de estadísticas",
					Points: []string{
						"Permite analizar resultados por día, semana, mes o año, mediante filtros de fecha dinámicos.",
						"Muestra KPIs principales como ventas totales, productos vendidos y ganancias del período seleccionado.",
						"Integra gráficos para visualizar la evolución de la recaudación y facilitar la toma de decisiones.",
					},
				},
				{
					Title: "Clientes",
					Image: "/static/images/projects/invoicing-system/clients-screen.webp",
					Alt:   "Gestión de clientes",
					Points: []string{
						"Permite registrar y administrar clientes, centralizando la información de contacto y relación comercial.",
						"Facilita la asociación de ventas a clientes, mejorando el seguimiento histórico de operaciones.",
						"Ofrece un listado estructurado para consulta y gestión rápida desde el sistema interno.",
					},
				},
				{
					Title: "Movimientos",
					Points: []string{
						"Presenta un historial de ventas/movimientos con filtros por rango temporal (últimas 24 h, última semana o todo el historial).",
						"Muestra la información en una tabla operativa con datos clave: cliente, producto, precio aplicado, estado y fecha/hora.",
						"Incluye paginación, manejo de estados (cargando, error, vacío) y navegación eficiente de grandes volúmenes de datos.",
					},
				},
				{
					Title: "Proveedores",
					Points: []string{
						"Permite gestionar proveedores, almacenando información clave para el abastecimiento de productos.",
						"Facilita la organización del origen del stock, vinculando productos con sus proveedores correspondientes.",
						"Centraliza la información para mejorar el control operativo y administrativo del negocio.",
					},
				},
				{
					Title: "Suscripciones",
					Points: []string{
						"Administra el concepto de planes y tenants, orientado a un sistema multi-usuario cerrado.",
						"Permite diferenciar niveles de acceso o funcionalidades según el plan contratado.",
						"Sienta la base para un modelo SaaS escalable, con control de usuarios y permisos.",
					},
				},
				{
					Title: "Cuenta / Ayuda",
					Points: []string{
						"Proporciona una sección de gestión de cuenta, accesible desde el navbar del sistema.",
						"Incluye accesos a ayuda o soporte, orientados a guiar al usuario dentro de la aplicación.",
						"Centraliza acciones sensibles como logout y configuración básica del usuario.",
					},
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
			DateRange: "2023 - 2024",
			Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},

			TitleES: "Creativistas Web",
			TitleEN: "Creativistas Web",
			DescES:  "Web de tests psicológicos (Big 5) con envío de resultados por email.",
			DescEN:  "Psychological tests (Big 5) with email delivery of results.",

			Screenshots: []string{
				"/static/images/projects/creativistas/creativistas-1.webp",
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
			DateRange: "2025 - Actualidad",
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
