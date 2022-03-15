package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"password"`
	IsAmbassador bool   `json:"is_ambassador"`
}

func (user *User) SetPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
	return hashedPassword
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
