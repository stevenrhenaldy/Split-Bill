package request

type RegisterRequest struct {
	Name              string `json:"name" validate:"required,min=2,max=100"`
	Username          string `json:"username" validate:"required,min=2,max=100"`
	Password          string `json:"password" validate:"required,min=6,max=100"`
	ConfirmPassword   string `json:"confirm_password" validate:"required,min=6,max=100,eqfield=Password"`
	DefaultCurrencyID string `json:"default_currency" validate:"required,min=3,max=3"`
	Email             string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=100,eqfield=Password"`
}
