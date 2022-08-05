package config

type Emitter struct {
	UserLoginStreamTopic     string `env:"USER_LOGIN_STREAM_TOPIC" envDefault:"user-login"`
	UserLoginStreamTopicNpar int    `env:"USER_LOGIN_STREAM_TOPIC_N_PAR" envDefault:"8"`
}
