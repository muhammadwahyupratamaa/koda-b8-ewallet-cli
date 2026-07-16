CREATE TABLE "users"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "full_name" VARCHAR(100) NOT NULL,
    "user_name" VARCHAR(40) NOT NULL UNIQUE,
    "age" INT NOT NULL,
    "address" TEXT NOT NULL,
    "email" VARCHAR(100) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW() 
);

