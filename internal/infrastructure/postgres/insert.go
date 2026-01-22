package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/ankitghl/activity-timeline-backend/internal/domain"
	"github.com/google/uuid"
)

const insertEventSQL = `
	INSERT INTO events (
		id, 
		kind,
		payload,
		event_time,
		recorded_at
	)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (id) DO NOTHING
	RETURNING recorded_at;
`

func (s *Store) InsertEvent(ctx context.Context, e domain.NewEvent, now time.Time) (*domain.AcceptedEvent, error) {
	id, err := uuid.Parse(e.ID)
	if err != nil {
		return nil, err
	}

	var recordedAt time.Time

	err = s.db.QueryRowContext(ctx, insertEventSQL, id, e.Kind, e.Payload, e.EvenTime, now).Scan(&recordedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &domain.AcceptedEvent{
		ID:          e.ID,
		RecorderdAt: recordedAt,
	}, nil
}
