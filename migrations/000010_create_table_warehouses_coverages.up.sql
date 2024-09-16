CREATE TABLE IF NOT EXISTS warehouses_coverages (
    id bigint NOT NULL,
    warehouse_id bigint NOT NULL,
    province VARCHAR(25) NOT NULL,
    city VARCHAR(50) NOT NULL,
    district VARCHAR(50) NOT NULL,
    sub_district VARCHAR(50) NOT NULL,
    zipcode VARCHAR(6) NOT NULL,
    tax NUMERIC(10, 3) NOT NULL,
    delivery_fee NUMERIC(10, 3) NOT NULL,
    service_fee NUMERIC(10, 3) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT warehouses_coverages_fk_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);