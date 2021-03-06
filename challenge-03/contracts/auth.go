package contracts

type AuthRequest struct {
	Username string
	Password string
}

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	ExpiresAt   int
}
