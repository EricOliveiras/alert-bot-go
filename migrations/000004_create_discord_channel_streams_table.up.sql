CREATE TABLE IF NOT EXISTS discord_channel_streams (
  id SERIAL PRIMARY KEY,
  discord_channel_id TEXT REFERENCES discord_channels(channel_id),
  stream_id BIGINT REFERENCES streams(id)
);