package strategies

import "context"

type IAccountStrategy interface {
	SwitchAccountRemoval(ctx context.Context, value []byte)
}
