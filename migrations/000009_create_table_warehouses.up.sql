CREATE TABLE IF NOT EXISTS warehouses (
    id bigint NOT NULL,  
    status VARCHAR(14) NOT NULL,      
    name VARCHAR(225) NOT NULL,  
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT warehouses_unique_name UNIQUE (name)
);