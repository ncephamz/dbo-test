CREATE TABLE IF NOT EXISTS customers_addresses (
    id bigint NOT NULL,
    customer_id bigint NOT NULL,
    province VARCHAR(25) NOT NULL,
    city VARCHAR(50) NOT NULL,
    district VARCHAR(50) NOT NULL,
    sub_district VARCHAR(50) NOT NULL,
    zipcode VARCHAR(6) NOT NULL,
    address TEXT NOT NULL,
    note TEXT NULL,
    google_map VARCHAR(225) NULL,
    is_main BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT customers_addresses_fk_customers FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);