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

func SingleResource(r *domain.Order) *orderResource {
	order := &orderResource{
		Code: r.Code,
		// TODO handle order resource
		CreatedAt: r.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: r.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	order.Id = r.ID
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
