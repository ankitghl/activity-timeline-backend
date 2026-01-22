package domain

import "time"

type NewEvent struct {
	ID       string
	Kind     EventKind
	Payload  []byte
	EvenTime time.Time
}

type AcceptedEvent struct {
	ID          string
	RecorderdAt time.Time
}
