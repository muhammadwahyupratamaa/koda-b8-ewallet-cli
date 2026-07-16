ALTER TABLE transactions
ALTER COLUMN sender_wallet_id SET NOT NULL;

ALTER TABLE transactions
ALTER COLUMN receiver_wallet_id SET NOT NULL;