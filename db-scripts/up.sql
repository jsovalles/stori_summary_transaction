DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions;

CREATE TABLE accounts (
    account_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    account_email TEXT NOT NULL UNIQUE
);

CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);

INSERT INTO accounts (account_id, account_email) VALUES('1c123230-5c31-4f39-836d-fe426bbb4d2a','storitest0@gmail.com');