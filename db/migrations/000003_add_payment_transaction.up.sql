CREATE TABLE payment_transactions (
    patrx_no VARCHAR(55) PRIMARY KEY,
    patrx_created_on date NOT NULL,
    patrx_debet DECIMAL(18,2) NOT NULL,
    patrx_credit DECIMAL(18,2) NOT NULL,
    patrx_notes VARCHAR(125),
    patrx_acctno_from VARCHAR(30) REFERENCES user_accounts(usac_account_no), 
    patrx_acctno_to VARCHAR(30) REFERENCES user_accounts(usac_account_no), 
    patrx_patrx_ref VARCHAR(55) REFERENCES payment_transactions(patrx_no), 
    patrx_traty_id INTEGER REFERENCES transaction_type(traty_id) 
);