package user

import (
	"alirah/app/domain"
)

type Resource struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func SingleResource(user *domain.User) *Resource {
	return &Resource{
		ID:       user.ID,
		FullName: user.FirstName + " " + user.LastName,
		Email:    user.Email,
	}
}

func Collection(users *[]domain.User) []*Resource {
	resources := make([]*Resource, 0)
	for _, value := range *users {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
