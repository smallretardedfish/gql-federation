CREATE TABLE IF NOT EXISTS medications
(
    id              BIGSERIAL PRIMARY KEY,
    code            JSONB,
    form            JSONB,
    manufacturer_id VARCHAR(50),
    patient_id      VARCHAR(50)
);