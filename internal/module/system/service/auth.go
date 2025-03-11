package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/smartmemos/memos/internal/config"
	"github.com/smartmemos/memos/internal/module/system/model"
)

func (s *Service) SignIn(ctx context.Context, req *model.SignInRequest) (user *model.User, err error) {
	user, err = s.dao.FindUser(ctx, &model.FindUserFilter{Username: req.Username})
	if err != nil {
		err = errors.Errorf("failed to find user by username %s", req.Username)
		return
	}
	if user.Status == model.Archived {
		err = errors.Errorf("user has been archived with username %s", req.Username)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received.
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		err = errors.New("unmatched email and password")
		return
	}
	expireTime := time.Now().Add(model.AccessTokenDuration)
	if req.NeverExpire {
		// Set the expire time to 100 years.
		expireTime = time.Now().Add(100 * 365 * 24 * time.Hour)
	}
	tokenStr, err := s.doSignIn(ctx, user, model.AccessTokenAudienceName, model.KeyID, expireTime)
	if err != nil {
		err = errors.Errorf("failed to sign in, err: %s", err)
		return
	}
	logrus.Infof("token: %s", tokenStr)
	// resp = &model.SignInResponse{
	// 	AccessToken: tokenStr,
	// 	ExpireTime:  expireTime,
	// }
	return
}

func (s *Service) doSignIn(ctx context.Context, user *model.User, audience, keyId string, expireTime time.Time) (tokenStr string, err error) {
	accessToken, err := s.GenerateAccessToken(ctx, user.ID, audience, keyId, expireTime)
	if err != nil {
		err = errors.Errorf("failed to generate tokens, err: %s", err)
		return
	}
	token := &model.AccessToken{
		Token:       accessToken.Token,
		Description: "user login",
	}
	logrus.Info(token)
	// if err = s.userService.UpsertAccessToken(ctx, user.ID, token); err != nil {
	// 	err = errors.Errorf("failed to upsert access token to store, err: %s", err)
	// 	return
	// }
	return accessToken.Token, nil
}

// GenerateAccessToken generates an access token.
func (s *Service) GenerateAccessToken(_ context.Context, userID int64, audience, keyId string, expirationTime time.Time) (accessToken *model.AccessToken, err error) {
	cfg := config.GetConfig().JWT

	issuedAt := time.Now()
	registeredClaims := jwt.RegisteredClaims{
		Issuer:   model.Issuer,
		Audience: jwt.ClaimStrings{audience},
		IssuedAt: jwt.NewNumericDate(issuedAt),
		Subject:  fmt.Sprint(userID),
	}
	if !expirationTime.IsZero() {
		registeredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	token.Header["kid"] = keyId

	// Create the JWT string.
	tokenStr, err := token.SignedString([]byte(cfg.Key))
	if err != nil {
		return
	}
	accessToken = &model.AccessToken{
		UserId:    userID,
		Token:     tokenStr,
		ExpiresAt: expirationTime,
		IssuedAt:  issuedAt,
	}
	return
}
