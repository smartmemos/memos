package service

import (
	"context"
	"time"

	"github.com/smartmemos/memos/internal/module/user/model"
)

func (s *Service) CreateAccessToken(ctx context.Context, req *model.CreateAccessTokenRequest) (token *model.AccessToken, err error) {
	tokenStr, err := s.authService.GenerateAccessToken(ctx, req.UserId, time.Now(), req.ExpiresAt)
	if err != nil {
		return
	}
	token = &model.AccessToken{
		UserId:      req.UserId,
		Token:       tokenStr,
		Description: req.Description,
		IssuedAt:    req.IssuedAt,
		ExpiresAt:   req.ExpiresAt,
	}
	err = s.dao.CreateAccessToken(ctx, token)
	return
}
