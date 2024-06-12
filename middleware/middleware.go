package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Jake-Andrews/go-web-dev/types"
)

func WithUser(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.URL.Path, "/public") {
            next.ServeHTTP(w, r)
            return
        }
        user := types.AuthenticatedUser{
            Email: "sneed@sneed.com",
            LoggedIn: true,
        }
        ctx := context.WithValue(r.Context(), types.UserReqKey, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
    return http.HandlerFunc(fn)
}
