package order

import (
	"alirah/app/domain"
)

type orderResource struct {
	Id        uint   `json:"id"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func SingleResource(p *domain.Order) *orderResource {
	order := &orderResource{
		Code:      p.Code,
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	order.Id = p.Id
	return order
}

func Collection(orders *[]domain.Order) []*orderResource {
	resources := make([]*orderResource, 0)
	for _, value := range *orders {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
