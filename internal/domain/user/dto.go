package user

type CreateUserDTO struct {
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}
