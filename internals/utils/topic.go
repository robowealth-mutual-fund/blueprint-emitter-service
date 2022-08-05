package utils

import (
	"log"

	"github.com/lovoo/goka"
)

// EnsureStreamExists is a convenience wrapper for TopicManager.EnsureStreamExists
func EnsureStreamExists(brokers []string, topic string, npar int) error {
	tm := createTopicManager(brokers, npar)
	defer tm.Close()
	err := tm.EnsureStreamExists(topic, npar)
	if err != nil {
		log.Printf("Error creating kafka stream topic %s: %v", topic, err)
		return err
	}
	return nil
}

// EnsureTableExists is a convenience wrapper for TopicManager.EnsureTableExists
func EnsureTableExists(brokers []string, topic string, npar int) error {
	tm := createTopicManager(brokers, npar)
	defer tm.Close()
	err := tm.EnsureTableExists(string(topic), npar)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
		return err
	}
	return nil
}

func createTopicManager(brokers []string, npar int) goka.TopicManager {
	tmc := goka.NewTopicManagerConfig()

	tmc.Table.Replication = npar
	tmc.Stream.Replication = npar

	tm, err := goka.NewTopicManager(brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	return tm
}
