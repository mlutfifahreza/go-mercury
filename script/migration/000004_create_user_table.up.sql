CREATE TABLE IF NOT EXISTS  "users"
(
    "username"   VARCHAR(64) PRIMARY KEY,
    "password_hash" VARCHAR(60) NOT NULL,
    "roles" TEXT
);