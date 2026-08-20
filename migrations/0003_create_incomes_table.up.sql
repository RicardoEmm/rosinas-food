CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS incomes (
    id UUID PRIMARY KEY,
    description VARCHAR(250),
    product_type VARCHAR(10) NOT NULL,
    quantity INT NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL,
    total NUMERIC(10,2) NOT NULL,
    customer_id UUID NOT NULL,
    status VARCHAR(15) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers
);
