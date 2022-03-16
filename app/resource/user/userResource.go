package user

import (
	"alirah/app/domain"
)

type userResource struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func SingleResource(user *domain.User) *userResource {
	return &userResource{
		ID:       user.Id,
		FullName: user.FirstName + " " + user.LastName,
		Email:    user.Email,
	}
}

func Collection(users *[]domain.User) []*userResource {
	resources := make([]*userResource, 0)
	for _, value := range *users {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
