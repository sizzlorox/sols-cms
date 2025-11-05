package services

import (
	"github.com/sizzlorox/sols-cms/api/v1/repositories"
	"github.com/sizzlorox/sols-cms/pkg/entities"
)

type IDomainService interface {
	GetDomainBySlug(slug string) (*entities.Domain, error)
	GetDomains(limit int, offset int) ([]*entities.Domain, error)
	CreateDomain(data entities.CreateDomainDTO) (interface{}, error)
	UpdateDomainBySlug(slug string, data entities.UpdateDomainDTO) (interface{}, error)
	DeleteDomainBySlug(slug string) error
}

type domainService struct {
	r repositories.IDomainRepository
}

func NewDomainService(r repositories.IDomainRepository) IDomainService {
	return &domainService{
		r,
	}
}

func (s *domainService) GetDomainBySlug(slug string) (*entities.Domain, error) {
	return s.r.GetDomainBySlug(slug)
}

func (s *domainService) GetDomains(limit int, offset int) ([]*entities.Domain, error) {
	return s.r.GetDomains(limit, offset)
}

// TODO: Update interface{} if repository is returning the entity
func (s *domainService) CreateDomain(data entities.CreateDomainDTO) (interface{}, error) {
	return s.r.CreateDomain(data)
}

// TODO: Update interface{} if repository is returning the entity
func (s *domainService) UpdateDomainBySlug(slug string, data entities.UpdateDomainDTO) (interface{}, error) {
	return s.r.UpdateDomainBySlug(slug, data)
}

func (s *domainService) DeleteDomainBySlug(slug string) error {
	return s.r.DeleteDomainBySlug(slug)
}
