package request

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
