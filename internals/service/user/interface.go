package user

import "context"

//go:generate mockery --name=Service
type Service interface {
	Emit(ctx context.Context, key string, msg interface{}) error
}
