CREATE TABLE IF NOT EXISTS discord_channels (
  id UUID PRIMARY KEY,
  name TEXT,
  channel_id TEXT UNIQUE,
  server_id TEXT UNIQUE,
  stream_limit INT,
  user_id UUID REFERENCES users(id),
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);