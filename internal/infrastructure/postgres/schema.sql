CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY_KEY,
    kind TEXT NOT NULL,
    payload JSONB NOT NULL,

    event_time TIMESTAMPTZ NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS events_ordering_idx
ON events (
    event_time DESC
    recorded_at DESC 
    id DESC
)