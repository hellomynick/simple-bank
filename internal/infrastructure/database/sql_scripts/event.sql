-- name: CreateEvent :one
INSERT INTO events (id,
                    aggregate_id,
                    aggregate_version,
                    type_name,
                    payload,
                    metadata,
                    occurred_at)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetEventsByAggregateID :many
SELECT *
FROM events
WHERE aggregate_id = $1
ORDER BY aggregate_version ASC;