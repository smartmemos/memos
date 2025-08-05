package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

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
	userID, sessionID, err := getUserInfoFromCookie(req)
	if err != nil {
		if isUnauthorizeAllowedMethod(req.URL.Path) {
			return nil, nil
		}
		return nil, err
	}
	user, err := a.memosService.GetUserByID(req.Context(), userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user")
	}
	return &utils.Info{
		UserID:    user.ID,
		SessionID: sessionID,
	}, nil
}

func getUserInfoFromCookie(req *http.Request) (userID int64, sessionID string, err error) {
	cookie, err := req.Cookie("memos.access-token")
	if err == http.ErrNoCookie {
		return
	} else if err != nil {
		err = errors.Wrap(err, "internal server error")
		return
	}
	parts := strings.SplitN(cookie.Value, "-", 2)
	if len(parts) != 2 {
		err = errors.New("invalid cookie value")
		return
	}
	userID, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		err = errors.Wrap(err, "no login")
		return
	}
	sessionID = parts[1]
	return
}
