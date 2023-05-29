CREATE TABLE IF NOT EXISTS conditions
(
    id         BIGSERIAL PRIMARY KEY,
    code       JSONB,
    category   JSONB,
    severity   JSONB,
    patient_id VARCHAR(50)
);