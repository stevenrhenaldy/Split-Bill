package response

type UserResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
