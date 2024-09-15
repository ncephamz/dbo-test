CREATE TABLE IF NOT EXISTS admins_roles (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    admin_id bigint NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT admins_roles_fk_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT admins_roles_fk_admin FOREIGN KEY (admin_id) REFERENCES admins (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);