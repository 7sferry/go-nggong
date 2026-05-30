CREATE TABLE IF NOT EXISTS dog_prices (
    id SERIAL PRIMARY KEY,
    dog_id INTEGER NOT NULL UNIQUE,
    price NUMERIC(19, 2) NOT NULL,
    price_type VARCHAR NOT NULL
);

INSERT INTO dog_prices (id, dog_id, price, price_type) VALUES (1, 1, 15000.50, 'purchase') ON CONFLICT (id) DO NOTHING;
INSERT INTO dog_prices (id, dog_id, price, price_type) VALUES (2, 2, 8250.75, 'adoption') ON CONFLICT (id) DO NOTHING;
INSERT INTO dog_prices (id, dog_id, price, price_type) VALUES (3, 3, 3420.00, 'stud') ON CONFLICT (id) DO NOTHING;

SELECT setval('dog_prices_id_seq', (SELECT MAX(id) FROM dog_prices));
