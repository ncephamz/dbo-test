CREATE TABLE IF NOT EXISTS roles_permissions (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    permission VARCHAR(225) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT roles_permissions_fk_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);