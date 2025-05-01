package request

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthRegisterRequest struct {
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required,min=6"`
	Firstname string `json:"first_name" form:"first_name" validate:"required,min=3"`
	Lastname  string `json:"last_name" form:"last_name"`
}
