package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/wl_super_backend_frontend/controllers"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", controllers.Home)
	app.Post("/api/alladmin", controllers.Adminhome)
	app.Post("/api/detailadmin", controllers.AdminDetail)
	app.Post("/api/saveadmin", controllers.AdminSave)

	app.Post("/api/alladminrule", controllers.Adminrulehome)
	app.Post("/api/saveadminrule", controllers.AdminruleSave)

	app.Post("/api/provider", controllers.Providerhome)
	app.Post("/api/providersave", controllers.ProviderSave)
	app.Post("/api/categame", controllers.Categamehome)
	app.Post("/api/categamesave", controllers.CategameSave)
	app.Post("/api/gamesave", controllers.GameSave)
	app.Post("/api/catebank", controllers.Catebankhome)
	app.Post("/api/catebanksave", controllers.CatebankSave)
	app.Post("/api/banktypesave", controllers.BankTypeSave)
	app.Post("/api/curr", controllers.Currencyhome)
	app.Post("/api/currsave", controllers.CurrencySave)
	app.Post("/api/master", controllers.Masterhome)
	app.Post("/api/masteragenadmin", controllers.Masteragenadmin)
	app.Post("/api/mastersave", controllers.MasterSave)
	app.Post("/api/masteradminsave", controllers.MasteradminSave)
	app.Post("/api/masteragensave", controllers.MasteragenSave)
	app.Post("/api/masteragenadminsave", controllers.MasteragenadminSave)

	return app
}
