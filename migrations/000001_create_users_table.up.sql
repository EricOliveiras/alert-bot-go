CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT,
    email TEXT UNIQUE,
    password TEXT,
    channel_limit INT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);