--Drop the foreign key constraints first
ALTER TABLE reviews DROP CONSTRAINT fk_product;
ALTER TABLE reviews DROP CONSTRAINT fk_user;

DROP TABLE IF EXISTS reviews;