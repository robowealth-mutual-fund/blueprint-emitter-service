package config

type Kafka struct {
	Brokers []string `env:"BROKERS" envDefault:"localhost:9094"`

	Emitter Emitter
}
