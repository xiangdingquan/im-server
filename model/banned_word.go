package model

import (
	"context"
)

func FixString(ctx context.Context, org string) string {
	return FixBannedWord(ctx, org)
}
