CREATE TABLE IF NOT EXISTS products (
  id CHAR(32) PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  description VARCHAR(64) NOT NULL,
  timestamp timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  price decimal(12,3) NOT NULL
);

SET timezone = 'Asia/Jakarta';