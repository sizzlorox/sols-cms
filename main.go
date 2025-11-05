package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/godruoyi/go-snowflake"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	v1routes "github.com/sizzlorox/sols-cms/api/v1/routes"
	"github.com/sizzlorox/sols-cms/pkg/providers/config"
	"github.com/sizzlorox/sols-cms/pkg/providers/database"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	config, err := config.NewConfigProvider()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		return
	}
	db, err := database.NewDatabaseProvider(config.GetDSN(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	snowflake.SetMachineID(uint16(config.MACHINE_ID))

	app := fiber.New()

	sessConfig := session.Config{
		Expiration:     30 * time.Minute,
		KeyLookup:      "cookie:__Host-session",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}
	store := session.New(sessConfig)
	csrfErrorHandler := func(c *fiber.Ctx, err error) error {
		fmt.Printf("CSRF Error: %v Request: %v From: %v\n", err, c.OriginalURL(), c.IP())
		switch c.Accepts("html", "json") {
		case "json":
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "403 Forbidden",
			})
		case "html":
			return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
				"Title":     "Error",
				"Error":     "403 Forbidden",
				"ErrorCode": "403",
			})
		default:
			return c.Status(fiber.StatusForbidden).SendString("403 Forbidden")
		}
	}

	app.Use(etag.New())
	app.Use(helmet.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    csrf.ConfigDefault.KeyGenerator(),
		Except: []string{"__Host-csrf", "__Host-session"},
	}))
	app.Use(csrf.New(csrf.Config{
		Session:        store,
		KeyLookup:      "header:" + csrf.HeaderName,
		CookieName:     "__Host-csrf",
		CookieSameSite: "Lax",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		ContextKey:     "csrf",
		ErrorHandler:   csrfErrorHandler,
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1routes.RegisterDomainRoutes(v1, db)

	go func() {
		if config.ENABLE_TLS {
			m := &autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(os.Getenv("APP_DOMAIN")),
				Cache:      autocert.DirCache("./certs"),
			}
			cfg := &tls.Config{
				GetCertificate: m.GetCertificate,
				// By default NextProtos contains the "h2"
				// This has to be removed since Fasthttp does not support HTTP/2
				// Or it will cause a flood of PRI method logs
				// http://webconcepts.info/concepts/http-method/PRI
				NextProtos: []string{
					"http/1.1", "acme-tls/1",
				},
			}
			ln, err := tls.Listen("tcp", ":443", cfg)
			if err != nil {
				panic(err)
			}

			// Start server
			log.Fatal(app.Listener(ln))
			return
		}
		log.Fatal(app.Listen(":3000"))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	_ = db.Close()

	fmt.Println("Fiber was successful shutdown.")
}
