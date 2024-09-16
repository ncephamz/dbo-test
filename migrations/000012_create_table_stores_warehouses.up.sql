CREATE TABLE IF NOT EXISTS stores_warehouses (
    id bigint NOT NULL,
    store_id bigint NOT NULL,
    warehouse_coverage_id bigint NOT NULL,
    product_id bigint NOT NULL,
    qty int NOT NULL,
    price NUMERIC(10, 3) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT stores_warehouses_fk_stores FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT stores_warehouses_fk_warehouses_coverages FOREIGN KEY (warehouse_coverage_id) REFERENCES warehouses_coverages (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT stores_warehouses_fk_products FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);