package view

import (
	"context"
	"strconv"

	"github.com/Jake-Andrews/go-web-dev/types"
)

func AuthenticatedUser(ctx context.Context) types.AuthenticatedUser {
	user, ok := ctx.Value(types.UserReqKey).(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{}
	}
	return user
}

func String(i int) string {
	return strconv.Itoa(i)
}
