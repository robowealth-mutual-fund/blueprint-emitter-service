package config

type Emitter struct {
	User UserEmitter
}

type TopicManager struct {
	NumStreamParReplicas int `env:"NUM_STREAM_PARTITION_REPLICAS" envDefault:"3"`
	NumTbParReplicas     int `env:"NUM_TABLE_PARTITION_REPLICAS" envDefault:"3"`
}

type UserEmitter struct {
	TopicStream  string `env:"TOPIC_STREAM" envDefault:"users"`
	NumPar       int    `env:"NUM_PAR" envDefault:"3"`
	TopicManager TopicManager
}
