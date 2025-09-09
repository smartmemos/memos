package service

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) CreateSession(ctx context.Context, req *model.CreateSessionRequest) (session *model.Session, err error) {
	user, err := s.dao.FindUser(ctx, &model.UserFilter{
		Username: db.Eq(req.Username),
	})
	if err != nil {
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received.
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		err = errors.Wrap(err, "invalid password")
		return
	}
	session = &model.Session{
		UserID:      user.ID,
		AccessToken: "",
	}
	if err = s.dao.CreateSession(ctx, session); err != nil {
		return
	}
	return
}
