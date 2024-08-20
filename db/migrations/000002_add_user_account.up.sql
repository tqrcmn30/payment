CREATE TABLE IF NOT EXISTS user_accounts (
    usac_id SERIAL PRIMARY KEY,
    usac_account_no VARCHAR(30) UNIQUE NOT NULL,
    usac_balance DECIMAL(18,2) NOT NULL,
    usac_created_on DATE NOT NULL DEFAULT CURRENT_DATE,
    usac_buty_id INTEGER,
    usac_user_id INTEGER
);