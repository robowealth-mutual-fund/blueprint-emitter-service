package user

import (
	"context"
)

func (i *Implement) Emit(ctx context.Context, key string, msg interface{}) error {
	err := i.emitter.UsrEmitter.EmitSync(key, msg)
	if err != nil {
		return err
	}

	return nil
}
