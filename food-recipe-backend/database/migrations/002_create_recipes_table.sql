CREATE TABLE recipes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    ingredients TEXT[],
    instructions TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
