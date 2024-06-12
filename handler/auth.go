package handler

import (
    "net/http"
    "github.com/Jake-Andrews/go-web-dev/view/auth"
)
func HandleSigninIndex(w http.ResponseWriter, r *http.Request) error {
    return auth.Signin().Render(r.Context(), w)
}

