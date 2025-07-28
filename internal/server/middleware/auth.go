package middleware

import (
	"context"
	"errors"
	"net/http"

	"connectrpc.com/authn"
	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/auth"
)

type Auth struct {
	authService auth.Service
}

func NewAuth(i do.Injector) *Auth {
	return &Auth{
		authService: do.MustInvoke[auth.Service](i),
	}
}

func (a *Auth) Auth(connectHandler http.Handler) http.Handler {
	middleware := authn.NewMiddleware(a.authenticate)
	handler := middleware.Wrap(connectHandler)
	return handler
}

func (a *Auth) authenticate(_ context.Context, req *http.Request) (any, error) {
	token, ok := authn.BearerToken(req)
	if !ok {
		if isUnauthorizeAllowedMethod(req.URL.Path) {
			return nil, nil
		}
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid authorization"))
	}
	userId, err := a.authService.Authenticate(req.Context(), token)
	if err != nil {
		if isUnauthorizeAllowedMethod(req.URL.Path) {
			return nil, nil
		}
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid authorization"))
	}
	authn.SetInfo(req.Context(), userId)

	// if isOnlyForAdminAllowedMethod(req.URL.Path) && user.Role != model.RoleHost && user.Role != model.RoleAdmin {
	// 	return nil, errors.Errorf("user %d is not admin", user.ID)
	// }

	return userId, nil
}
