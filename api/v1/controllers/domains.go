package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sizzlorox/sols-cms/api/v1/services"
	"github.com/sizzlorox/sols-cms/pkg/entities"
)

type IDomainController interface {
	GetDomainBySlug(ctx *fiber.Ctx) error
	GetDomains(ctx *fiber.Ctx) error
	CreateDomain(ctx *fiber.Ctx) error
	UpdateDomainBySlug(ctx *fiber.Ctx) error
	DeleteDomainBySlug(ctx *fiber.Ctx) error
}

type domainController struct {
	s services.IDomainService
}

func NewDomainController(s services.IDomainService) IDomainController {
	return &domainController{
		s,
	}
}

func (dc *domainController) GetDomainBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	if slug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid domain slug"})
	}
	domain, err := dc.s.GetDomainBySlug(slug)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(domain)
}

func (dc *domainController) GetDomains(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid limit"})
	}
	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid offset"})
	}
	domains, err := dc.s.GetDomains(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(domains)
}

func (dc *domainController) CreateDomain(ctx *fiber.Ctx) error {
	var data entities.CreateDomainDTO
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	domain, err := dc.s.CreateDomain(data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(domain)
}

func (dc *domainController) UpdateDomainBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	if slug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid domain slug"})
	}
	data := entities.UpdateDomainDTO{}
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	domain, err := dc.s.UpdateDomainBySlug(slug, data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(domain)
}

func (dc *domainController) DeleteDomainBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	if slug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid domain slug"})
	}
	if err := dc.s.DeleteDomainBySlug(slug); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
