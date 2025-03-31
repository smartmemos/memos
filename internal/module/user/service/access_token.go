package service

import (
	"context"
	"slices"
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

func (s *Service) ListAccessTokens(ctx context.Context, req *model.ListAccessTokensRequest) (list []*model.AccessToken, err error) {
	filter := &model.FindAccessTokenFilter{}
	tokens, err := s.dao.FindAccessTokens(ctx, filter)
	if err != nil {
		return
	}
	for _, token := range tokens {
		item, aErr := s.authService.Authenticate(ctx, token.Token)
		if aErr != nil {
			continue
		}

		item.ID = token.ID
		item.Description = token.Description
		list = append(list, item)
	}
	// Sort by issued time in descending order.
	slices.SortFunc(list, func(i, j *model.AccessToken) int {
		return int(i.IssuedAt.Sub(j.IssuedAt))
	})
	return
}
