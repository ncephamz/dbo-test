CREATE TABLE IF NOT EXISTS orders (
    id bigint NOT NULL,
    customer_address_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    status VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT orders_fk_customers_addresses FOREIGN KEY (customer_address_id) REFERENCES customers_addresses (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT orders_fk_customers FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);