package handler

import (
	"log/slog"
	"net/http"

	"github.com/Jake-Andrews/go-web-dev/types"
)

func getAuthenticatedUser(r *http.Request) types.AuthenticatedUser {
    user, ok := r.Context().Value(types.UserReqKey).(types.AuthenticatedUser)
    if !ok {
        return types.AuthenticatedUser{}
    }
    return user
}

func MakeHandler(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        if err := h(w,r); err!=nil {
            slog.Error("MakeHandler", "err", err, "path", r.URL.Path)
        }
    }
}
