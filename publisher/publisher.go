package publisher

import "context"

type Publisher interface {
	PublishMessage(ctx context.Context, msg interface{}) error
}
