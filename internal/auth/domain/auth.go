package domain

type Request struct {
	Email   string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=25"`
}

type Response struct {
	Token string `json:"token"`
}