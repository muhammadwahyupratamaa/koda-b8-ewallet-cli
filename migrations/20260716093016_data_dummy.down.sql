DELETE FROM wallets
WHERE wallet_number IN (
    'WAL000001',
    'WAL000002',
    'WAL000003'
);

DELETE FROM users
WHERE user_name IN (
    'wahyu',
    'budi',
    'siti'
);