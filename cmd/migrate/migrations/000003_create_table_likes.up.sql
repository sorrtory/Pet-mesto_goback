CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    card_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) on delete set null on update cascade,
    FOREIGN KEY (card_id) REFERENCES cards (id) on delete cascade on update cascade
)
