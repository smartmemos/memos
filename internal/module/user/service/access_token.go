package service

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
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

func (s *Service) DeleteAccessToken(ctx context.Context, req *model.DeleteAccessTokenRequest) (err error) {
	userID, err := grpc_util.GetUserID(ctx)
	if err != nil {
		return
	}
	if req.UserID != userID {
		err = errors.New("permission denied")
		return
	}
	filter := &model.FindAccessTokenFilter{
		Token: req.AccessToken,
	}
	err = s.dao.DeleteAccessToken(ctx, filter)
	return
}
