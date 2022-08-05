package user

import (
	"github.com/lovoo/goka"
	"github.com/robowealth-mutual-fund/blueprint-emitter-service/internals/config"
)

type Emitter struct {
	*goka.Emitter
}

func NewUserEmitter(codec goka.Codec) (*Emitter, error) {
	conf := config.NewConfiguration()
	tmc := goka.NewTopicManagerConfig()

	nb := len(conf.Kafka.Brokers)
	tmc.Table.Replication = nb
	tmc.Stream.Replication = nb

	tm, err := goka.NewTopicManager(conf.Kafka.Brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		return nil, err
	}

	err = tm.EnsureStreamExists(conf.Emitter.User.TopicStream, conf.Emitter.User.NumPar)
	if err != nil {
		return nil, err
	}

	emitter, err := goka.NewEmitter(conf.Kafka.Brokers, goka.Stream(conf.Emitter.User.TopicStream), codec)
	if err != nil {
		return nil, err
	}

	return &Emitter{
		Emitter: emitter,
	}, nil
}
