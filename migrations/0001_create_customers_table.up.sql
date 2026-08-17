CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(75) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

CREATE UNIQUE INDEX idx_customers_phone ON customers (phone);