CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(25) NOT NULL,
    in_game_name VARCHAR(25) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP default now()
);