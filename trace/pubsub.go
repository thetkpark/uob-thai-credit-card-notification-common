package trace

import (
	"github.com/google/uuid"
)

func GetCorrelationIdFromPubSubAttributes(attrs map[string]string) string {
	if attrs == nil {
		return uuid.NewString()
	}
	if attrs[CorrelationIdKey] == "" {
		return uuid.NewString()
	}
	return attrs[CorrelationIdKey]
}

func AttachCorrelationIdToPubSubAttributes(attrs map[string]string, correlationId string) map[string]string {
	if attrs == nil {
		attrs = make(map[string]string)
	}
	attrs[CorrelationIdKey] = correlationId
	return attrs
}
