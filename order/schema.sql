CREATE TABLE IF NOT EXISTS orders (
  id CHAR(32) PRIMARY KEY,
  account_id CHAR(32) NOT NULL,
  total_price decimal(12,3) NOT NULL
  timestamp timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS order_products (
  order_id CHAR(32) REFERENCES orders (id) ON DELETE CASCADE,
  product_id CHAR(32),
  quantity INT NOT NULL,
  PRIMARY KEY (product_id, order_id)
);
