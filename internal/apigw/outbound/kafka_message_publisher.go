package outbound

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"vc/internal/apigw/apiv1"
	"vc/pkg/logger"
	"vc/pkg/messagebroker/kafka"
	"vc/pkg/model"
	"vc/pkg/trace"

	"github.com/IBM/sarama"
)

type kafkaMessageProducer struct {
	client *kafka.SyncProducerClient
}

// New creates a new instance of a kafka event publisher used by apigw
func New(ctx context.Context, cfg *model.Cfg, tracer *trace.Tracer, log *logger.Log) (apiv1.EventPublisher, error) {
	saramaConfig := kafka.CommonProducerConfig(cfg)
	client, err := kafka.NewSyncProducerClient(ctx, saramaConfig, cfg, tracer, log.New("kafka_message_producer_client"))
	if err != nil {
		return nil, err
	}

	return &kafkaMessageProducer{client: client}, nil

}

// Upload publish a UploadRequest message to a Kafka topic
func (s *kafkaMessageProducer) Upload(uploadRequest *apiv1.UploadRequest) error {
	if uploadRequest == nil {
		return errors.New("param uploadRequest is nil")
	}

	jsonMarshaled, err := json.Marshal(uploadRequest)
	if err != nil {
		return err
	}

	//TODO(mk): make header code below including in other kafka publisher generic and move to kafka client
	paramType := reflect.TypeOf(uploadRequest).Elem().Name()
	typeHeader := []byte(paramType)
	headers := []sarama.RecordHeader{
		{Key: []byte(kafka.TypeOfStructInMessageValue), Value: typeHeader},
	}

	return s.client.PublishMessage(kafka.TopicUpload, uploadRequest.Meta.DocumentID, jsonMarshaled, headers)
}

// Close closes all resources used/started by the publisher
func (s *kafkaMessageProducer) Close(ctx context.Context) error {
	if s.client != nil {
		return s.client.Close(ctx)
	}
	return nil
}
