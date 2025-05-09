-- Revert the updated_by column in customer_car table back to VARCHAR(5)
ALTER TABLE customer_car
ALTER COLUMN updated_by TYPE VARCHAR(5);