CREATE TABLE products
(
    id                SERIAL PRIMARY KEY,
    NAME              TEXT NOT NULL,
    category_id       INTEGER NOT NULL, --foreign key refers to the category table
    long_description  TEXT NOT NULL, -- stored as plain html
    short_description TEXT NOT NULL, -- stored as plain html
    thumbnail_image   TEXT NOT NULL,
    carousel_images   TEXT NOT NULL,
    slug              TEXT NOT NULL UNIQUE,
    price             INTEGER NOT NULL,
    discount          INTEGER DEFAULT 0,
    inventory         SMALLINT DEFAULT 0,
    sku               TEXT NOT NULL,
    status            BOOLEAN DEFAULT true,
    weight            DOUBLE PRECISION DEFAULT 0,
    width             DOUBLE PRECISION DEFAULT 0,
    length            DOUBLE PRECISION DEFAULT 0,
    height            DOUBLE PRECISION DEFAULT 0,
    total_rating      DOUBLE PRECISION DEFAULT 5,
    amount_sold       INTEGER DEFAULT 0,
    created_at        TIMESTAMP DEFAULT Now(),
    updated_at        TIMESTAMP DEFAULT Now(),
    -- Foreign key constraint
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
        ON DELETE SET NULL
);