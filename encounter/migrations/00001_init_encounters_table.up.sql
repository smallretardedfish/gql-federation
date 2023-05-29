CREATE TABLE IF NOT EXISTS encounters
(
    id         BIGSERIAL PRIMARY KEY,
    status     VARCHAR(25),
    type       JSONB,
    period     JSONB,
    patient_id BIGINT
);