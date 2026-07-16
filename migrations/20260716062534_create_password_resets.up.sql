CREATE TABLE "password_resets"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" BIGINT NOT NULL UNIQUE,
    "token" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "expired_at" TIMESTAMP NOT NULL,

    CONSTRAINT fk_password_resets_user
                FOREIGN KEY (user_id)
                REFERENCES users(id)
);