package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"log"
)

type PubSubConfig struct {
	ProjectID string `env:"PUBSUB_PROJECT_ID"`
	TopicID   string `env:"PUBSUB_TOPIC_ID"`
}

type PubSubPublisher struct {
	topic *pubsub.Topic
}

func NewPubSubPublisher(projectId, topicName string) *PubSubPublisher {
	client, err := pubsub.NewClient(context.Background(), projectId)
	if err != nil {
		log.Fatalln(err)
	}
	topic := client.Topic(topicName)
	return &PubSubPublisher{topic: topic}
}

func (p PubSubPublisher) PublishMessage(ctx context.Context, msg interface{}, attrs map[string]string) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	result := p.topic.Publish(ctx, &pubsub.Message{Data: b, Attributes: attrs})
	if _, err = result.Get(ctx); err != nil {
		return err
	}
	return nil
}
