package kafkastreams

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/kafka_stream/emitter/user"
)

type Emitter struct {
	UsrEmitter *user.Emitter
}

func NewEmitter(usrLoginEmitter *user.Emitter) *Emitter {
	return &Emitter{
		UsrEmitter: usrLoginEmitter,
	}
}
