package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		ID:           "",
		Email:        dto.Email,
		Username:     dto.Username,
		PasswordHash: "",
	}
}
func (u *User) GeneratePasswordHash(password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hashedPassword)
	// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)

	return nil
}

type CreateUserDTO struct {
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
