
INSERT INTO users (
    full_name,
    user_name,
    age,
    address,
    email,
    password
) VALUES
(
    'Muhammad Wahyu Pratama',
    'wahyu',
    23,
    'Baturaja',
    'wahyu@mail.com',
    '123456'
),
(
    'Budi Santoso',
    'budi',
    25,
    'Jakarta',
    'budi@mail.com',
    '123456'
),
(
    'Siti Aisyah',
    'siti',
    22,
    'Bandung',
    'siti@mail.com',
    '123456'
);

INSERT INTO wallets (
    user_id,
    wallet_number,
    balance,
    status
) VALUES
(
    1,
    'WAL000001',
    1000000,
    'ACTIVE'
),
(
    2,
    'WAL000002',
    500000,
    'ACTIVE'
),
(
    3,
    'WAL000003',
    250000,
    'ACTIVE'
);