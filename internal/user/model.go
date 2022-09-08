package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	UserName     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
