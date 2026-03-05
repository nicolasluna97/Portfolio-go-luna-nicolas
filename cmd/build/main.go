package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Ajustá si cambiás rutas
const (
	tplDir    = "web/templates"
	staticDir = "web/static"
	outDir    = "public_html"
)

func main() {
	// 1) limpiar salida
	if err := os.RemoveAll(outDir); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatal(err)
	}

	// 2) copiar /static
	if err := copyDir(staticDir, filepath.Join(outDir, "static")); err != nil {
		log.Fatalf("copiando static: %v", err)
	}

	// 3) render páginas
	// Home -> /index.html
	if err := renderPage(
		filepath.Join(tplDir, "layouts", "base.html"),
		filepath.Join(tplDir, "pages", "home.html"),
		filepath.Join(outDir, "index.html"),
		homeData(),
	); err != nil {
		log.Fatalf("render home: %v", err)
	}

	// Proyectos -> /projects/<slug>/index.html
	pages := []struct {
		slug     string
		template string
		data     any
	}{
		{"invoicing-system", "project-invoicing.html", projectInvoicingData()},
		{"creativistas-web", "project-creativistas.html", projectCreativistasData()},
		{"tienda-nube", "project-tiendanube.html", projectTiendaNubeData()},
	}

	for _, p := range pages {
		out := filepath.Join(outDir, "projects", p.slug, "index.html")
		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			log.Fatal(err)
		}

		if err := renderPage(
			filepath.Join(tplDir, "layouts", "base.html"),
			filepath.Join(tplDir, "pages", p.template),
			out,
			p.data,
		); err != nil {
			log.Fatalf("render %s: %v", p.slug, err)
		}
	}

	log.Printf("✅ Build OK -> %s", outDir)
}

// -------------------- render --------------------

func renderPage(layoutPath, pagePath, outputPath string, data any) error {
	tpl := template.Must(template.ParseFiles(layoutPath, pagePath))

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "base", data); err != nil {
		return err
	}

	// si necesitás rewrites para hosting estático, acá es donde se ajusta.
	html := buf.Bytes()

	return os.WriteFile(outputPath, html, 0644)
}

// -------------------- copy helpers --------------------

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

// -------------------- data (copiado de tu handlers.go) --------------------
// IMPORTANTE: esto replica tus structs para poder compilar el build
// (sin importar el package internal). Si después querés, lo hacemos más prolijo.

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

func homeData() PageData {
	return PageData{
		Title:   "Portfolio - Luna Nicolás",
		Lang:    "es",
		Role:    "Desarrollador Full Stack",
		Name:    "Luna Nicolás Ezequiel",
		HideNav: false,
		Social: SocialLinks{
			Github:   "https://github.com/nicolasluna97",
			Email:    "mailto:nicolassluna1997@gmail.com",
			Linkedin: "https://www.linkedin.com/in/nicolas-lunaok/",
		},
		StackPrimary: []StackItem{
			{Key: "angular", Label: "Angular"},
			{Key: "nestjs", Label: "NestJS"},
			{Key: "postgresql", Label: "PostgreSQL"},
			{Key: "docker", Label: "Docker"},
			{Key: "github", Label: "GitHub"},
			{Key: "hostinger", Label: "Hostinger"},
		},
		StackOther: []StackItem{
			{Key: "go", Label: "Go"},
			{Key: "javascript", Label: "JavaScript"},
			{Key: "mysql", Label: "MySQL"},
			{Key: "claude", Label: "Claude AI"},
			{Key: "html5", Label: "HTML5"},
			{Key: "css3", Label: "CSS3"},
			{Key: "VirtualBox", Label: "VirtualBox"},
			{Key: "postman", Label: "Postman"},
		},
		Projects: []ProjectCard{
			{
				URL:       "/projects/invoicing-system",
				Thumb:     "/static/images/projects/invoicing-system/thumb.webp",
				DateRange: "",
				Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},
				TitleES:   "Sistema de Facturación (Multi-tenant)",
				TitleEN:   "Invoicing System (Multi-tenant)",
				DescES:    "Plataforma de facturación e inventario con roles, clientes, productos y movimientos de stock.",
				DescEN:    "Multi-tenant invoicing and inventory platform with roles, customers, products, and stock movement tracking.",
			},
			{
				URL:       "/projects/creativistas-web",
				Thumb:     "/static/images/projects/creativistas/creativistas-thumb.webp",
				DateRange: "",
				Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},
				TitleES:   "Creativistas Web",
				TitleEN:   "Creativistas Web",
				DescES:    "Web de tests psicológicos, con envío de resultados por email.",
				DescEN:    "Psychological tests, with automated email delivery of results.",
			},
			{
				URL:       "/projects/tienda-nube",
				Thumb:     "/static/images/projects/tiendanube/thumb-tiendanube.webp",
				DateRange: "",
				Tech:      []string{"TiendaNube", "UI/UX", "CSS", "HTML"},
				TitleES:   "Tienda Online (TiendaNube)",
				TitleEN:   "Online Store (TiendaNube)",
				DescES:    "E-commerce en TiendaNube: personalización del theme, configuración de catálogo/variantes y optimización del flujo de compra para autogestión.",
				DescEN:    "TiendaNube e-commerce: theme customization, catalog/variants setup, and checkout flow optimization for client self-management.",
			},
		},
	}
}

func projectInvoicingData() ProjectPageData {
	return ProjectPageData{
		Title:   "Invoicing System - Luna Nicolás",
		Lang:    "es",
		Role:    "Desarrollador Full Stack",
		Name:    "Luna Nicolás Ezequiel",
		HideNav: true,
		Social: SocialLinks{
			Github:   "https://github.com/nicolasluna97",
			Email:    "mailto:nicolassluna1997@gmail.com",
			Linkedin: "https://www.linkedin.com/in/nicolas-lunaok/",
		},
		Slug:      "invoicing-system",
		HeroImage: "/static/images/projects/invoicing-system/thumb.webp",
		DateRange: "",
		Tech:      []string{"Angular", "NestJS", "PostgreSQL", "Docker"},
		TitleES:   "Sistema de Facturación (Multi-tenant)",
		TitleEN:   "Invoicing System (Multi-tenant)",
		DescES:    "Plataforma de facturación e inventario con roles, clientes, productos, movimientos de stock y estadísticas para empresas o autonomos. La página web posse actulaización de stock en tiempo real además modo oscuro y claro. Actualmente en desarrollo en beta 2.0, cuando se sumen las funcionalidades faltantes, se subirá a producción.",
		DescEN:    "Multi-tenant invoicing and inventory platform with roles, customers, products, stock movements, and analytics for businesses and self-employed professionals. It provides real-time stock updates and supports both dark and light themes. Currently in beta 2.0; once the remaining features are completed, it will be released to production.",
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
					{ES: "Permite crear productos mediante formularios modales, asignando categorías, stock inicial, precio de compra y múltiples precios de venta por producto.",
						EN: "Allows creating products via modal forms, assigning categories, initial stock, purchase cost, and multiple selling prices per product."},
					{ES: "Habilita edición masiva de productos seleccionados, con validaciones por campo y feedback visual de errores.",
						EN: "Enables bulk editing of selected products, with per-field validations and visual error feedback."},
					{ES: "Soporta eliminación segura de uno o varios productos mediante modales de confirmación para evitar acciones accidentales.",
						EN: "Supports safe deletion of one or multiple products through confirmation modals to prevent accidental actions."},
				},
			},
			{
				TitleES: "Ventas",
				TitleEN: "Sales",
				Image:   "/static/images/projects/invoicing-system/sales-screen.webp",
				AltES:   "Flujo de ventas",
				AltEN:   "Sales flow",
				Points: []ProjectPoint{
					{ES: "Ofrece un flujo de venta rápido, pensado para uso operativo, donde se seleccionan productos, cantidades y lista de precios.",
						EN: "Provides a fast, operations-oriented sales flow where products, quantities, and a price list are selected."},
					{ES: "Soporta múltiples precios por producto, permitiendo calcular totales según el precio aplicado en cada venta.",
						EN: "Supports multiple prices per product, calculating totals based on the price applied to each sale."},
					{ES: "Al confirmar la operación, descuenta automáticamente el stock, impactando en inventario en tiempo real.",
						EN: "On confirmation, it automatically decreases stock, updating inventory in real time."},
				},
			},
			{
				TitleES: "Estadísticas",
				TitleEN: "Analytics",
				Image:   "/static/images/projects/invoicing-system/statistics-screen.webp",
				AltES:   "Panel de estadísticas",
				AltEN:   "Analytics dashboard",
				Points: []ProjectPoint{
					{ES: "Permite analizar resultados por día, semana, mes o año, mediante filtros de fecha dinámicos.",
						EN: "Allows analyzing results by day, week, month, or year using dynamic date filters."},
					{ES: "Muestra KPIs principales como ventas totales, productos vendidos y ganancias del período seleccionado.",
						EN: "Shows key KPIs such as total sales, items sold, and profit for the selected period."},
					{ES: "Integra gráficos para visualizar la evolución de la recaudación y facilitar la toma de decisiones.",
						EN: "Includes charts to visualize revenue trends and support decision-making."},
				},
			},
			{
				TitleES: "Clientes",
				TitleEN: "Customers",
				Image:   "/static/images/projects/invoicing-system/clients-screen.webp",
				AltES:   "Gestión de clientes",
				AltEN:   "Customer management",
				Points: []ProjectPoint{
					{ES: "Permite registrar y administrar clientes, centralizando la información de contacto y relación comercial.",
						EN: "Lets you register and manage customers, centralizing contact details and business relationship data."},
					{ES: "Facilita la asociación de ventas a clientes, mejorando el seguimiento histórico de operaciones.",
						EN: "Enables linking sales to customers, improving historical tracking of operations."},
					{ES: "Ofrece un listado estructurado para consulta y gestión rápida desde el sistema interno.",
						EN: "Provides a structured list for quick lookup and management inside the system."},
				},
			},
		},
	}
}

func projectCreativistasData() ProjectPageData {
	return ProjectPageData{
		Title:   "Creativistas Web - Luna Nicolás",
		Lang:    "es",
		Role:    "Desarrollador Full Stack",
		Name:    "Luna Nicolás Ezequiel",
		HideNav: true,
		Social: SocialLinks{
			Github:   "https://github.com/nicolasluna97",
			Email:    "mailto:nicolassluna1997@gmail.com",
			Linkedin: "https://www.linkedin.com/in/nicolas-lunaok/",
		},
		Slug:      "creativistas",
		HeroImage: "/static/images/projects/creativistas/creativistas-thumb.webp",
		DateRange: "",
		Tech:      []string{"Next.js", "React", "MongoDB", "SendGrid"},
		TitleES:   "Creativistas Web",
		TitleEN:   "Creativistas Web",
		DescES:    "Web de tests psicológicos, con envío de resultados por email.",
		DescEN:    "Psychological tests, with email delivery of results.",
		Screenshots: []string{
			"/static/images/projects/creativistas/creativistas-2.webp",
		},
	}
}

func projectTiendaNubeData() ProjectPageData {
	return ProjectPageData{
		Title:   "TiendaNube - Luna Nicolás",
		Lang:    "es",
		Role:    "Desarrollador Full Stack",
		Name:    "Luna Nicolás Ezequiel",
		HideNav: true,
		Social: SocialLinks{
			Github:   "https://github.com/nicolasluna97",
			Email:    "mailto:nicolassluna1997@gmail.com",
			Linkedin: "https://www.linkedin.com/in/nicolas-lunaok/",
		},
		Slug:      "tienda-nube",
		HeroImage: "/static/images/projects/tiendanube/tiendanube-website.webp",
		DateRange: "",
		Tech:      []string{"TiendaNube", "UI/UX", "HTML", "CSS"},
		TitleES:   "Tienda Online (TiendaNube)",
		TitleEN:   "Online Store (TiendaNube)",
		DescES:    "Implementación y personalización de tienda en TiendaNube.",
		DescEN:    "Implementation and customization of a TiendaNube store.",
	}
}

// solo para evitar warning si querés loggear algo rápido
func _unused(args ...any) { fmt.Println(args...) }
