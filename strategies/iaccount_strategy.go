package strategies

import "context"

//go:generate mockgen -destination=./mocks/iaccount_strategy.go -package=mocks --build_flags=--mod=mod . IAccountStrategy
type IAccountStrategy interface {
	SwitchAccountRemoval(ctx context.Context, value []byte)
}
