package user

import (
	kafkastreams "github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/repository/kafka_stream"
)

type Implement struct {
	emitter *kafkastreams.Emitter
}

func NewService(emitter *kafkastreams.Emitter) Service {
	return &Implement{
		emitter: emitter,
	}
}
