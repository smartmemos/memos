package model

import (
	"time"
)

const (
	// issuer is the issuer of the jwt token.
	Issuer = "memos"
	// Signing key section. For now, this is only used for signing, not for verifying since we only
	// have 1 version. But it will be used to maintain backward compatibility if we change the signing mechanism.
	KeyID = "v1"
	// AccessTokenAudienceName is the audience name of the access token.
	AccessTokenAudienceName = "user.access-token"
	AccessTokenDuration     = 7 * 24 * time.Hour
)

type SignInRequest struct {
	Username    string
	Password    string
	NeverExpire bool
}

const (
	TableAccessToken = "access_token"
)
