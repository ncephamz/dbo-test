CREATE TABLE IF NOT EXISTS admins (
    id bigint NOT NULL,  
    username VARCHAR(14) NOT NULL,      
    password TEXT NOT NULL,       
    name VARCHAR(50) NULL,  
    email VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT admins_unique_username UNIQUE (username)
);