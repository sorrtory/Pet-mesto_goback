CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    card_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (card_id) REFERENCES cards (id)
)
