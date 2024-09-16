CREATE TABLE IF NOT EXISTS stores (
    id bigint NOT NULL,  
    level VARCHAR(14) NOT NULL,      
    name VARCHAR(225) NOT NULL,  
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT stores_unique_name UNIQUE (name)
);