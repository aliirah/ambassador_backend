package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"password"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func (u *User) SetPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = hashedPassword
	return hashedPassword
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
