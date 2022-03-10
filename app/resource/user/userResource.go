package user

import "alirah/app/domain"

type userResource struct {
	ID           uint   `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func SingleResource(user *domain.User) *userResource {
	return &userResource{
		ID:           user.Id,
		FullName:     user.FirstName + " " + user.LastName,
		Email:        user.Email,
		IsAmbassador: user.IsAmbassador,
	}
}
