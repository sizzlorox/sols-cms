package v1routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sizzlorox/sols-cms/api/v1/controllers"
	"github.com/sizzlorox/sols-cms/api/v1/repositories"
	"github.com/sizzlorox/sols-cms/api/v1/services"
	"github.com/sizzlorox/sols-cms/pkg/providers/database"
)

func RegisterDomainRoutes(r fiber.Router, db database.IRepository) {
	domain := r.Group("/domains")

	domainRepository := repositories.NewDomainRepository(db)
	domainService := services.NewDomainService(domainRepository)
	domainController := controllers.NewDomainController(domainService)

	domain.Get("/:domainSlug", domainController.GetDomainBySlug)

	domain.Get("/", domainController.GetDomains)
	domain.Post("/", domainController.CreateDomain)
	domain.Put("/:domainSlug", domainController.UpdateDomainBySlug)
	domain.Delete("/:domainSlug", domainController.DeleteDomainBySlug)
}
