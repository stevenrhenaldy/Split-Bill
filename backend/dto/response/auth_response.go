package response

type AuthResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type MeResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Username          string `json:"username"`
	EmailVerifiedAt   string `json:"email_verified_at"`
	DefaultCurrencyID string `json:"default_currency_id"`
}
