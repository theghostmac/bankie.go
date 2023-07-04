-- List contents of account table.
SELECT * FROM account;

-- List all existing tables.
SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public'
ORDER BY table_name;

-- Insert value into account table.
INSERT INTO account (first_name, last_name, email, bank_number, balance, created_at)
VALUES ('Ghost', 'Mac', 'macbobbychibuzor@gmail.com', 1234567890, 100.50, NOW());

-- Check type of `balance` column in account table.
SELECT column_name, data_type
FROM information_schema.columns
WHERE table_name = 'account' AND column_name = 'balance';

-- Change type
ALTER TABLE account ALTER COLUMN balance TYPE float8;

-- Delete a row
DELETE FROM account WHERE id = 2;
