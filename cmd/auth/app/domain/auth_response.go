package domain

type LoginResponse struct {
	TokenType    string `json:"token_type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int32  `json:"expires"`
}
