package link

import (
	"alirah/app/domain"
)

type linkResource struct {
	Id        uint   `json:"id"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func SingleResource(p *domain.Link) *linkResource {
	link := &linkResource{
		Code:      p.Code,
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	link.Id = p.ID
	return link
}

func Collection(links *[]domain.Link) []*linkResource {
	resources := make([]*linkResource, 0)
	for _, value := range *links {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
