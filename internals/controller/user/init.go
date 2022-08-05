package user

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/service/user/wrapper"
)

type Controller struct {
	service *wrapper.Wrapper
}

func NewController(s *wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
