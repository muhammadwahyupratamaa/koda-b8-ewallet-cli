CREATE TABLE "transactions"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "sender_wallet_id" BIGINT,
    "receiver_wallet_id" BIGINT, 
    "amount" DECIMAL NOT NULL ,
    "type" VARCHAR(100) NOT NULL,
    "status" VARCHAR(100) NOT NULL DEFAULT 'ACTIVE',
    "description" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_wallet_sender
                FOREIGN KEY(sender_wallet_id)
                REFERENCES wallets(id),

    CONSTRAINT fk_wallet_receiver
                FOREIGN KEY(receiver_wallet_id)
                REFERENCES wallets(id)
);