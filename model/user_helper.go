package model

import (
	"context"
)

type UserHelper interface {
	GetMutableUsers(ctx context.Context, idList ...int32) MutableUsers
}
