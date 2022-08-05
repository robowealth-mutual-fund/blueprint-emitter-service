package kafkastreams

import (
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/emitter/user_login_emitter"
)

type Emitter struct {
	UsrLoginEmitter *userloginemitter.Emitter
}

func NewEmitter(usrLoginEmitter *userloginemitter.Emitter) *Emitter {
	return &Emitter{UsrLoginEmitter: usrLoginEmitter}
}
