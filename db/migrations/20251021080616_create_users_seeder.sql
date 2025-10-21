-- migrate:up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users (name, email, password, phone, meta, verified_at)
VALUES
  ('John Doe', 'john@example.com', crypt('password123', gen_salt('bf')), '081234567890', null, NOW());

-- migrate:down
DELETE FROM users WHERE email = 'john@example.com';
