CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username STRING (32) NOT NULL UNIQUE,
    display_name STRING (40),
    icon STRING,
    INDEX(username)
);
