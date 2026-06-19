CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,

    username VARCHAR(50) UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    mfa_enabled BOOLEAN DEFAULT FALSE,
    mfa_secret TEXT,

    failed_attempts INTEGER DEFAULT 0,
    locked_until TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMPTZ
);