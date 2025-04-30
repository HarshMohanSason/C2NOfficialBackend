CREATE TABLE categories
(
    id                   SERIAL PRIMARY KEY,
    NAME                 TEXT NOT NULL UNIQUE,
    size_chart           TEXT NOT NULL,
    how_to_measure_image TEXT NOT NULL,
    customization_pdf    TEXT DEFAULT NULL,
    created_at           TIMESTAMP DEFAULT Now(),
    updated_at           TIMESTAMP DEFAULT Now()
);