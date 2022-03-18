package product

import (
	"alirah/app/domain"
)

type productResource struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func SingleResource(p *domain.Product) *productResource {
	product := &productResource{
		Title:       p.Title,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	product.Id = p.ID
	return product
}

func Collection(products *[]domain.Product) []*productResource {
	resources := make([]*productResource, 0)
	for _, value := range *products {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
