package repositories

import (
	"context"
	"simplebank/internal/infrastructure/database/generated"
)

type EventStore interface {
	generated.Queries
	TransferTx(ctx context.Context, arg generated.CreateEventParams)
}
