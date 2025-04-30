CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    NAME       VARCHAR(255) NOT NULL,
    email      TEXT NOT NULL UNIQUE,
    password   TEXT,
    auth_type  VARCHAR(20) NOT NULL,
    is_admin   BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);