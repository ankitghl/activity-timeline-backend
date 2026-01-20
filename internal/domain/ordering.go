package domain

const (
	OrderingSQL = `
		ORDER BY
			event_time DESC,
			recorderd_at DESC,
			id DESC
	`
)