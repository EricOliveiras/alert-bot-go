CREATE TABLE IF NOT EXISTS streams (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE,
  image_url TEXT,
  is_live BOOLEAN,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);