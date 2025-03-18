ALTER TABLE recipes 
ADD CONSTRAINT fk_recipes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE payments 
ADD CONSTRAINT fk_payments_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
ADD CONSTRAINT fk_payments_recipe FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE;
