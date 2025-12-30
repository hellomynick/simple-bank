CREATE TABLE events
(
    id                UUID PRIMARY KEY,
    aggregate_id      UUID        NOT NULL,
    aggregate_version BIGINT      NOT NULL,
    type_name         VARCHAR     NOT NULL,
    payload           JSONB       NOT NULL,
    occurred_at       TIMESTAMPTZ NOT NULL,
    metadata          JSONB       NOT NULL DEFAULT '{}'::jsonb,
    CONSTRAINT uk_aggregate_version UNIQUE (aggregate_id, aggregate_version)
)
