-- DB gallery_db

CREATE TABLE "products"
(
    "id"          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "title"       VARCHAR(256) NOT NULL,
    "image_url"   VARCHAR(256) NOT NULL,
    "description" TEXT         NOT NULL
);

CREATE TABLE "stores"
(
    "id"   VARCHAR(64) PRIMARY KEY,
    "name" VARCHAR(256) NOT NULL,
    "icon" VARCHAR(256) NOT NULL
);

CREATE TABLE "links"
(
    "id"         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "product_id" BIGINT NOT NULL,
    "store_id"   VARCHAR(64) NOT NULL,
    "link"       VARCHAR(256) NOT NULL
);
ALTER TABLE "links"
    ADD CONSTRAINT "links_product_id_foreign" FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "links"
    ADD CONSTRAINT "links_store_id_foreign" FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
