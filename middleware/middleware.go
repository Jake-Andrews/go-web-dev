package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Jake-Andrews/go-web-dev/types"
)

type UserKey int
const (
    UserReqKey UserKey = iota
)

func WithUser(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.URL.Path, "/public") {
            next.ServeHTTP(w, r)
            return
        }
        user := types.AuthenticatedUser{}
        ctx := context.WithValue(r.Context(), UserReqKey, user)
        next.ServeHTTP(w, r.WithContext(ctx))

    }
    return http.HandlerFunc(fn)
}
