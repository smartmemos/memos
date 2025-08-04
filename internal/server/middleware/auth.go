package middleware

import (
	"context"
	"net/http"
	"strconv"

	"connectrpc.com/authn"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/pkg/utils"
)

type Auth struct {
	memosService memos.Service
}

func NewAuth(i do.Injector) *Auth {
	return &Auth{
		memosService: do.MustInvoke[memos.Service](i),
	}
}

func (a *Auth) Auth(connectHandler http.Handler) http.Handler {
	middleware := authn.NewMiddleware(a.authenticate)
	return middleware.Wrap(connectHandler)
}

func (a *Auth) authenticate(_ context.Context, req *http.Request) (any, error) {
	cookie, err := req.Cookie("memos.access-token")
	if err != nil {
		if isUnauthorizeAllowedMethod(req.URL.Path) {
			return nil, nil
		}
		if err == http.ErrNoCookie {
			return nil, err
		}
		return nil, errors.Wrap(err, "internal server error")
	}
	userID, err := strconv.ParseInt(cookie.Value, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "no login")
	}
	user, err := a.memosService.GetUserByID(req.Context(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}

	return &utils.Info{
		UserID:    user.ID,
		SessionID: "",
	}, nil
}
