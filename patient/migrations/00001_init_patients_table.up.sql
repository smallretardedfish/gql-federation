CREATE TABLE IF NOT EXISTS patients
(
  id         BIGSERIAL PRIMARY KEY,
  gender     VARCHAR(50),
  birth_date VARCHAR(100),
  name       JSONB,
  address    JSONB,
  telecom    JSONB
);

-- INSERT INTO patients (id, gender, birth_date, name, address, telecom)
-- VALUES
-- ('Male', '1990-01-01', '{"use": "official", "family": "Smith", "given": ["John"]}'::jsonb, '[{"use": "home", "line": ["123 Main St"], "city": "New York", "state": "NY", "postalCode": "12345", "country": "USA"}]'::jsonb, '[{"value": "john@example.com", "system": "email", "use": "home"}, {"value": "555-123-4567", "system": "phone", "use": "mobile"}]'::jsonb),
-- ('Female', '1985-03-15', '{"use": "official", "family": "Johnson", "given": ["Emily"]}'::jsonb, '[{"use": "home", "line": ["456 Elm St"], "city": "Los Angeles", "state": "CA", "postalCode": "98765", "country": "USA"}]'::jsonb, '[{"value": "emily@example.com", "system": "email", "use": "home"}, {"value": "555-987-6543", "system": "phone", "use": "mobile"}]'::jsonb),
-- ('Male', '1978-07-22', '{"use": "official", "family": "Brown", "given": ["Michael"]}'::jsonb, '[{"use": "home", "line": ["789 Oak St"], "city": "Chicago", "state": "IL", "postalCode": "54321", "country": "USA"}]'::jsonb, '[{"value": "michael@example.com", "system": "email", "use": "home"}, {"value": "555-567-8901", "system": "phone", "use": "mobile"}]'::jsonb),
-- ('Female', '1992-09-10', '{"use": "official", "family": "Davis", "given": ["Emma"]}'::jsonb, '[{"use": "home", "line": ["789 Maple Ave"], "city": "San Francisco", "state": "CA", "postalCode": "54321", "country": "USA"}]'::jsonb, '[{"value": "emma@example.com", "system": "email", "use": "home"}, {"value": "555-789-0123", "system": "phone", "use": "mobile"}]'::jsonb),
-- ( 'Male', '1982-05-03', '{"use": "official", "family": "Wilson", "given": ["James"]}'::jsonb, '[{"use": "home", "line": ["321 Oak Ave"], "city": "Seattle", "state": "WA", "postalCode": "98765", "country": "USA"}]'::jsonb, '[{"value": "james@example.com", "system": "email", "use": "home"}, {"value": "555-012-3456", "system": "phone", "use": "mobile"}]'::jsonb),
-- ( 'Female', '1995-11-27', '{"use": "official", "family": "Anderson", "given": ["Olivia"]}'::jsonb, '[{"use": "home", "line": ["456 Elm St"], "city": "Boston", "state": "MA", "postalCode": "23456", "country": "USA"}]'::jsonb, '[{"value": "olivia@example.com", "system": "email", "use": "home"}, {"value": "555-345-6789", "system": "phone", "use": "mobile"}]'::jsonb),
-- ( 'Male', '1987-12-08', '{"use": "official", "family": "Turner", "given": ["Daniel"]}'::jsonb, '[{"use": "home", "line": ["789 Pine St"], "city": "Dallas", "state": "TX", "postalCode": "65432", "country": "USA"}]'::jsonb, '[{"value": "daniel@example.com", "system": "email", "use": "home"}, {"value": "555-987-6543", "system": "phone", "use": "mobile"}]'::jsonb);