CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    tag VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
