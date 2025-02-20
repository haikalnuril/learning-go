package dto

type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Gender               string `json:"gender"`
}

type RegisterResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}