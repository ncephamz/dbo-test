CREATE TABLE IF NOT EXISTS products (
    id bigint NOT NULL,
    store_id bigint NOT NULL,
    product_name VARCHAR(225) NOT NULL,
    image VARCHAR(225) NULL,
    uom VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT products_fk_stores FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);