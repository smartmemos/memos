package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/smartmemos/memos/internal/config"
	"github.com/smartmemos/memos/internal/module/auth/model"
	usermd "github.com/smartmemos/memos/internal/module/user/model"
)

func (s *Service) SignIn(ctx context.Context, req *model.SignInRequest) (accessToken *model.AccessToken, err error) {
	user, err := s.userService.GetUserByUsername(ctx, req.Username)
	if err != nil {
		err = errors.Errorf("failed to find user by username %s", req.Username)
		return
	}
	if user.Status == usermd.Archived {
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
		// Set the expire time to 10 years.
		expireTime = time.Now().Add(10 * 365 * 24 * time.Hour)
	}
	issuedAt := time.Now()
	token, err := s.generateAccessToken(ctx, user.ID, issuedAt, expireTime)
	if err != nil {
		return
	}
	accessToken = &model.AccessToken{
		UserId:      user.ID,
		Token:       token,
		Description: "login",
		IssuedAt:    issuedAt,
		ExpiresAt:   expireTime,
	}
	err = s.dao.CreateAccessToken(ctx, accessToken)
	return
}

func (s *Service) generateAccessToken(_ context.Context, userID int64, issuedAt, expirationTime time.Time) (tokenStr string, err error) {
	cfg := config.GetConfig().JWT

	registeredClaims := jwt.RegisteredClaims{
		Issuer:   model.Issuer,
		Audience: jwt.ClaimStrings{model.AccessTokenAudienceName},
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
	tokenStr, err = token.SignedString([]byte(cfg.Key))
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
		accessToken.IssuedAt = claims.IssuedAt.Time
	}
	if claims.ExpiresAt != nil {
		accessToken.ExpiresAt = claims.ExpiresAt.Time
	}
	return
}

func (s *Service) DeleteAccessToken(ctx context.Context, userId int64, token string) error {
	return s.dao.DeleteAccessToken(ctx, &model.FindAccessTokenFilter{
		UserId: userId,
		Token:  token,
	})
}

func (s *Service) ValidateAccessToken(ctx context.Context, userId int64, token string) (bool, error) {
	total, err := s.dao.CountAccessTokens(ctx, &model.FindAccessTokenFilter{
		UserId: userId,
		Token:  token,
	})
	return total > 0, err
}
