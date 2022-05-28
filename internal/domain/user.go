package domain

import (
	"fmt"
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

	return nil
}

func (u *User) CompareHashAndPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return fmt.Errorf("passwords don't match")
	}

	return nil
}

type CreateUserDTO struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type SignInUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
