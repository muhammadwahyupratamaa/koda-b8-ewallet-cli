ALTER TABLE transactions
ALTER COLUMN sender_wallet_id DROP NOT NULL;

ALTER TABLE transactions
ALTER COLUMN receiver_wallet_id DROP NOT NULL;