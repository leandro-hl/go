CREATE TABLE IF NOT EXISTS threads
(
    id          UUID PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts
(
    id        UUID PRIMARY KEY,
    thread_id UUID NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    title     TEXT NOT NULL,
    content   TEXT NOT NULL,
    votes     INT  NOT NULL
);

CREATE TABLE IF NOT EXISTS comments
(
    id      UUID PRIMARY KEY,
    post_id UUID NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    votes   INT  NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id       UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions
(
    token  TEXT PRIMARY KEY,
    data   BYTEA       NOT NULL,
    expiry TIMESTAMPTZ NOT NULL

);

CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry)