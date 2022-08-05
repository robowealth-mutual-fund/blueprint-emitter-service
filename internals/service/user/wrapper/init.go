package wrapper

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/service/user"
)

type Wrapper struct {
	Service user.Service
}

func NewWrapper(service user.Service) *Wrapper {
	return &Wrapper{
		Service: service,
	}
}
