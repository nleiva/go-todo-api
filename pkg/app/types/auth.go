package types

type LoginDTOBody struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=100"`
}

type LoginDTO struct {
	Account LoginDTOBody `json:"account"`
}

type RegisterDTOBody struct {
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required,min=6,max=100"`
	Firstname       string `json:"firstname" form:"firstname" validate:"omitempty,min=2"`
	Lastname        string `json:"lastname" form:"lastname" validate:"omitempty,min=2"`
}

type RegisterDTO struct {
	Account RegisterDTOBody `json:"account"`
}

type AuthResponseBody struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	Auth AuthResponseBody `json:"auth"`
}
