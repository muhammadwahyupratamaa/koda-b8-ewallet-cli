CREATE TABLE "wallets"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" BIGINT NOT NULL UNIQUE,
    "wallet_number" VARCHAR(100) NOT NULL UNIQUE,
    "balance" DECIMAL NOT NULL DEFAULT 0 ,
    "status" VARCHAR(100) NOT NULL ,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_wallets_user
                FOREIGN KEY(user_id)
                REFERENCES users(id)
);