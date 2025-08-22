package request

type UpdateUserRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	Email           string `json:"email" validate:"required,email"`
	DefaultCurrency string `json:"default_currency" validate:"required,min=3,max=3"`
}

type ChangeUserPasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=6,max=100"`
	NewPassword     string `json:"new_password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=100,eqfield=NewPassword"`
}

type ForgetUserPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetUserPasswordRequest struct {
	Password        string `json:"password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=100,eqfield=Password"`
}

type DeleteUserRequest struct {
	Password string `json:"password" validate:"required"`
}
