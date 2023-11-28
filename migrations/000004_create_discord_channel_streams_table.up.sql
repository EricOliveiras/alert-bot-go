CREATE TABLE IF NOT EXISTS discord_channel_streams (
  id SERIAL PRIMARY KEY,
  discord_channel_id UUID REFERENCES discord_channels(id),
  stream_id BIGINT REFERENCES streams(id)
);