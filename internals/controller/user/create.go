package user

import (
	"context"

	"github.com/opentracing/opentracing-go"

	apiV1 "github.com/robowealth-mutual-fund/blueprint-emitter-service/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.UserCreateRequest) (*apiV1.UserCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.user.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	err := c.service.Emit(ctx, "user-login", "")
	if err != nil {
		return nil, err
	}
	return &apiV1.UserCreateResponse{Id: int32(0)}, nil
}
