CREATE TABLE IF NOT EXISTS streams (
  id INT PRIMARY KEY,
  name TEXT UNIQUE,
  image_url TEXT,
  is_live BOOLEAN,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);