package trace

import (
	"context"
	"github.com/google/uuid"
	"github.com/thetkpark/uob-thai-credit-card-notification-common/logger"
	"log/slog"
)

const (
	CorrelationIdKey = "x-correlation-id"
)

func AddCorrelationIdToLogContext(ctx context.Context, correlationId string) context.Context {
	return logger.AppendCtxValue(ctx, slog.String(CorrelationIdKey, correlationId))
}

func GenerateCorrelationId() string {
	return uuid.NewString()
}
