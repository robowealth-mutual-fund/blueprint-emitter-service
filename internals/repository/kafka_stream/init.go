package kafkastreams

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/emitter/user"
)

type Emitter struct {
	UsrLoginEmitter *user.Emitter
}

func NewEmitter(usrLoginEmitter *user.Emitter) *Emitter {
	return &Emitter{
		UsrLoginEmitter: usrLoginEmitter,
	}
}
