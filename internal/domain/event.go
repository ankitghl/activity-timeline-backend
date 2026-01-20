package domain

import (
	"time"

	"github.com/google/uuid"
)

type EventKind string

type Event struct {
	ID         uuid.UUID
	Kind       EventKind
	Payload    []byte
	EventTime  time.Time
	RecordedAt time.Time
}
