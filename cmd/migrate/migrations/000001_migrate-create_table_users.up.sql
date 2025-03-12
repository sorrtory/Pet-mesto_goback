CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    password UUID UNIQUE,
    name VARCHAR(30) NOT NULL,
    about VARCHAR(100),
    avatar TEXT,
    cohort VARCHAR(20)
);
