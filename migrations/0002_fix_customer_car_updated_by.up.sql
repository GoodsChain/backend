-- Fix the updated_by column in customer_car table (was too small at VARCHAR(5))
ALTER TABLE customer_car
ALTER COLUMN updated_by TYPE VARCHAR(50);