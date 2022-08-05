package userloginemitter

import (
	"github.com/lovoo/goka"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
)

type Emitter struct {
	*goka.Emitter
}

func NewUserLoginEmitter() (*Emitter, error) {
	conf := config.NewConfiguration()

	tmc := goka.NewTopicManagerConfig()
	nb := len(conf.Kafka.Brokers)
	tmc.Table.Replication = nb
	tmc.Stream.Replication = nb

	tm, err := goka.NewTopicManager(conf.Kafka.Brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		return nil, err
	}

	err = tm.EnsureStreamExists(conf.Emitter.UserLoginStreamTopic, conf.Emitter.UserLoginStreamTopicNpar)
	if err != nil {
		return nil, err
	}

	emitter, err := goka.NewEmitter(conf.Kafka.Brokers, goka.Stream(conf.Emitter.UserLoginStreamTopic), nil)
	if err != nil {
		return nil, err
	}

	return &Emitter{
		Emitter: emitter,
	}, nil
}
