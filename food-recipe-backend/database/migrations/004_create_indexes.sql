CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_recipes_user_id ON recipes(user_id);
CREATE INDEX idx_payments_status ON payments(status);
