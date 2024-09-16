CREATE TABLE IF NOT EXISTS orders_details (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    store_warehouse_id bigint NOT NULL,
    qty int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT orders_details_fk_orders FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT orders_details_fk_stores_warehouses FOREIGN KEY (store_warehouse_id) REFERENCES stores_warehouses (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);