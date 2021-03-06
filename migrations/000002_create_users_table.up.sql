CREATE TABLE IF NOT EXISTS chalky.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username STRING (32) NOT NULL UNIQUE,
    display_name STRING (40),
    icon STRING,
    password_hash BYTES NOT NULL,
    INDEX(username)
);
