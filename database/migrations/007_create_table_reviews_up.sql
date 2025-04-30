CREATE TABLE reviews
(
    id         SERIAL PRIMARY KEY,
    COMMENT    TEXT NOT NULL,
    rating     SMALLINT NOT NULL CHECK (rating >= 0 AND rating <= 5),
    user_id    INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    images     TEXT[], -- Array of paths for images
    created_at TIMESTAMP DEFAULT Now(),
    updated_at TIMESTAMP DEFAULT Now(),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON
        DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE
        CASCADE
);