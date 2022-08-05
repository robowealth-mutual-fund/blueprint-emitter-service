package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func (w *Wrapper) Emit(ctx context.Context, key string, msg interface{}) error {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.User.Emit")
	defer sp.Finish()

	sp.LogKV("key", key)
	sp.LogKV("msg", msg)

	err := w.Service.Emit(ctx, key, msg)
	if err != nil {
		return err
	}

	sp.LogKV("err", err)

	return err
}
