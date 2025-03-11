package model

type SignInRequest struct {
	Username    string
	Password    string
	NeverExpire bool
}

type AccessToken struct {
	UserId      int64  `json:"user_id"`
	Token       string `json:"token"`
	Description string `json:"description"`
	IssuedAt    int64  `json:"issued_at"`
	ExpiresAt   int64  `json:"expires_at"`
}
