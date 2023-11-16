CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  discord_id TEXT UNIQUE,
  username TEXT,
  email TEXT UNIQUE,
  avatar TEXT,
  channel_limit INT,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);