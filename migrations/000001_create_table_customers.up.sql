CREATE TABLE IF NOT EXISTS customers (
    id bigint NOT NULL,  
    phone_number VARCHAR(14) NOT NULL,      
    email VARCHAR(100) NOT NULL,       
    password TEXT NOT NULL,       
    name VARCHAR(50) NULL,
    photo_profile VARCHAR(225) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT customers_unique_phone_number UNIQUE (phone_number)
);