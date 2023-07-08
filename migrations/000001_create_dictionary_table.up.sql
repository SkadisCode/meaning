-- migrate:up
CREATE TABLE IF NOT EXISTS dictionary (
    id SERIAL PRIMARY KEY,
    word VARCHAR(255) NOT NULL,
    entry TEXT NOT NULL,
    type VARCHAR(255),
    tier VARCHAR(255),
    semantic VARCHAR(255)
    abec VARCHAR(255)
);
