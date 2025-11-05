package repositories

import (
	"fmt"

	"github.com/sizzlorox/sols-cms/pkg/entities"
	"github.com/sizzlorox/sols-cms/pkg/providers/database"
)

type IDomainRepository interface {
	GetDomainBySlug(slug string) (*entities.Domain, error)
	GetDomains(limit int, offset int) ([]*entities.Domain, error)
	CreateDomain(data entities.CreateDomainDTO) (interface{}, error)
	UpdateDomainBySlug(slug string, data entities.UpdateDomainDTO) (interface{}, error)
	DeleteDomainBySlug(slug string) error
}

type domainRepository struct {
	db database.IRepository
}

func NewDomainRepository(db database.IRepository) IDomainRepository {
	return &domainRepository{
		db: db,
	}
}

func (dr *domainRepository) GetDomainBySlug(slug string) (*entities.Domain, error) {
	result := entities.Domain{
		Slug: slug,
	}
	dr.db.FindOne(&result)
	if result.Name == "" {
		return nil, fmt.Errorf("domain not found")
	}
	return &result, nil
}

func (dr *domainRepository) GetDomains(limit int, offset int) ([]*entities.Domain, error) {
	var results []entities.Domain
	dr.db.FindMany(&results, limit, offset)
	result := make([]*entities.Domain, len(results))
	for i, v := range results {
		result[i] = &v
	}
	return result, nil
}

func (dr *domainRepository) CreateDomain(data entities.CreateDomainDTO) (interface{}, error) {
	var results entities.Domain
	results.Name = data.Name
	results.Slug = data.Slug
	err := dr.db.InsertOne(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (dr *domainRepository) UpdateDomainBySlug(slug string, data entities.UpdateDomainDTO) (interface{}, error) {
	result := entities.Domain{
		Slug: slug,
	}
	if data.Name != nil {
		result.Name = *data.Name
	}
	if data.Slug != nil {
		result.Slug = *data.Slug
	}
	err := dr.db.UpdateOne(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dr *domainRepository) DeleteDomainBySlug(slug string) error {
	result := entities.Domain{
		Slug: slug,
	}
	err := dr.db.DeleteOne(&result)
	return err
}
