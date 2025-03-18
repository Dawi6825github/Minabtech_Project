CREATE OR REPLACE FUNCTION get_user_by_email(user_email TEXT)
RETURNS TABLE(id INT, full_name TEXT, email TEXT, role TEXT, created_at TIMESTAMP) AS $$
BEGIN
    RETURN QUERY 
    SELECT id, full_name, email, role, created_at
    FROM users
    WHERE email = user_email;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_user_purchases(user_id INT)
RETURNS TABLE(recipe_id INT, title TEXT, amount DECIMAL, status TEXT) AS $$
BEGIN
    RETURN QUERY 
    SELECT p.recipe_id, r.title, p.amount, p.status
    FROM payments p
    JOIN recipes r ON p.recipe_id = r.id
    WHERE p.user_id = user_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_recipes_by_user(user_id INT)
RETURNS TABLE(id INT, title TEXT, description TEXT, created_at TIMESTAMP) AS $$
BEGIN
    RETURN QUERY 
    SELECT id, title, description, created_at
    FROM recipes
    WHERE user_id = user_id;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_recipe_sales(recipe_id INT)
RETURNS INT AS $$
DECLARE total_sales INT;
BEGIN
    SELECT COUNT(*) INTO total_sales FROM payments WHERE recipe_id = recipe_id AND status = 'completed';
    RETURN total_sales;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION mark_payment_completed(tx_id TEXT)
RETURNS VOID AS $$
BEGIN
    UPDATE payments
    SET status = 'completed'
    WHERE transaction_id = tx_id;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_total_revenue()
RETURNS DECIMAL AS $$
DECLARE total DECIMAL;
BEGIN
    SELECT SUM(amount) INTO total FROM payments WHERE status = 'completed';
    RETURN COALESCE(total, 0);
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_user_count()
RETURNS INT AS $$
DECLARE count_users INT;
BEGIN
    SELECT COUNT(*) INTO count_users FROM users;
    RETURN count_users;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_active_users()
RETURNS TABLE(user_id INT, full_name TEXT, email TEXT) AS $$
BEGIN
    RETURN QUERY 
    SELECT u.id, u.full_name, u.email
    FROM users u
    JOIN payments p ON u.id = p.user_id
    WHERE p.created_at >= NOW() - INTERVAL '30 days'
    GROUP BY u.id;
END;
$$ LANGUAGE plpgsql;
