CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL DEFAULT 0
);
INSERT INTO products (name, price, stock)
VALUES ('Product 1', 19.99, 5),
    ('Product 2', 29.99, 10);