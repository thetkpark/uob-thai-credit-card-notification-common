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
