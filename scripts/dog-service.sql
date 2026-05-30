CREATE TABLE IF NOT EXISTS dogs (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    role VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS dog_contacts (
    id SERIAL PRIMARY KEY,
    dog_id INTEGER NOT NULL UNIQUE,
    phone VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    emergency VARCHAR NOT NULL
);

INSERT INTO dogs (id, name, email, role) VALUES (1, 'Rex', 'rex@example.com', 'guard') ON CONFLICT (id) DO NOTHING;
INSERT INTO dogs (id, name, email, role) VALUES (2, 'Bella', 'bella@example.com', 'companion') ON CONFLICT (id) DO NOTHING;
INSERT INTO dogs (id, name, email, role) VALUES (3, 'Max', 'max@example.com', 'hunter') ON CONFLICT (id) DO NOTHING;

INSERT INTO dog_contacts (id, dog_id, phone, address, emergency) VALUES (1, 1, '+1-555-0101', '123 Main St, New York', '+1-555-0901') ON CONFLICT (id) DO NOTHING;
INSERT INTO dog_contacts (id, dog_id, phone, address, emergency) VALUES (2, 2, '+1-555-0102', '456 Oak Ave, Chicago', '+1-555-0902') ON CONFLICT (id) DO NOTHING;
INSERT INTO dog_contacts (id, dog_id, phone, address, emergency) VALUES (3, 3, '+1-555-0103', '789 Pine Rd, Seattle', '+1-555-0903') ON CONFLICT (id) DO NOTHING;

SELECT setval('dogs_id_seq', (SELECT MAX(id) FROM dogs));
SELECT setval('dog_contacts_id_seq', (SELECT MAX(id) FROM dog_contacts));
