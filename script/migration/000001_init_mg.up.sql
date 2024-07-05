-- DB gallery_db

CREATE TABLE IF NOT EXISTS "products"
(
    "id"          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "title"       VARCHAR(256) NOT NULL,
    "image_url"   VARCHAR(256) NOT NULL,
    "description" TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS  "stores"
(
    "id"   VARCHAR(64) PRIMARY KEY,
    "name" VARCHAR(256) NOT NULL,
    "icon" VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS  "links"
(
    "id"         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "product_id" BIGINT NOT NULL REFERENCES products("id"),
    "store_id"   VARCHAR(64) NOT NULL REFERENCES stores("id"),
    "link"       VARCHAR(256) NOT NULL
);
