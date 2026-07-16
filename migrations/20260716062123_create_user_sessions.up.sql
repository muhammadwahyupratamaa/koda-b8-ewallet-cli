CREATE TABLE "user_sessions"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" BIGINT NOT NULL UNIQUE,
    "session_token" VARCHAR(255) NOT NULL UNIQUE,
    "expired_at" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_user_sessions_user
               FOREIGN KEY (user_id)
               REFERENCES users(id)
);