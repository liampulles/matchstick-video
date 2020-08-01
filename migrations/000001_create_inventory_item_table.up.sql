CREATE TABLE IF NOT EXISTS inventory_item(
   id SERIAL PRIMARY KEY,
   name VARCHAR(511) UNIQUE NOT NULL,
   location VARCHAR(255) UNIQUE NOT NULL,
   available BOOLEAN NOT NULL
);