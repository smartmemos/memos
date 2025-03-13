package service

import (
	"context"
	"fmt"
	"strconv"
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
	tokenStr, err := s.doSignIn(ctx, user, model.AccessTokenAudienceName, expireTime)
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

func (s *Service) doSignIn(ctx context.Context, user *model.User, audience string, expireTime time.Time) (tokenStr string, err error) {
	accessToken, err := s.GenerateAccessToken(ctx, user.ID, audience, expireTime)
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
func (s *Service) GenerateAccessToken(_ context.Context, userID int64, audience string, expirationTime time.Time) (accessToken *model.AccessToken, err error) {
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
	token.Header["kid"] = model.KeyID

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

func (in *Service) Authenticate(ctx context.Context, tokenStr string) (accessToken *model.AccessToken, err error) {
	if tokenStr == "" {
		err = errors.New("access token not found")
		return
	}
	cfg := config.GetConfig().JWT
	claims := &jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == model.KeyID {
				return []byte(cfg.Key), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		err = errors.New("Invalid or expired access token")
		return
	}
	userId, err := strconv.ParseInt(claims.Subject, 10, 32)
	if err != nil {
		return
	}
	accessToken = &model.AccessToken{
		UserId: userId,
		Token:  tokenStr,
	}
	if claims.IssuedAt != nil {
		// accessToken.IssuedAt = claims.IssuedAt.Unix()
	}
	if claims.ExpiresAt != nil {
		// accessToken.ExpiresAt = claims.ExpiresAt.Unix()
	}
	return
}
