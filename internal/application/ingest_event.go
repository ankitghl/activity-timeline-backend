package application

import (
	"context"
	"time"

	"github.com/ankitghl/activity-timeline-backend/internal/domain"
)

type EventWriter interface {
	InsertEvent(ctx context.Context, e domain.NewEvent, now time.Time) (*domain.AcceptedEvent, error)
}

type IngestEvents struct {
	writer EventWriter
	clock  func() time.Time
}

func NewIngestEvents(writer EventWriter, clock func() time.Time) *IngestEvents {
	return &IngestEvents{writer: writer, clock: clock}
}

func (uc *IngestEvents) Execute(ctx context.Context, events []domain.NewEvent) ([]domain.AcceptedEvent, error) {
	now := uc.clock()
	accpted := make([]domain.AcceptedEvent, 0, len(events))

	for _, e := range events {
		result, err := uc.writer.InsertEvent(ctx, e, now)
		if err != nil {
			return nil, err
		}

		if result != nil {
			accpted = append(accpted, *result)
		}
	}
	return accpted, nil
}
