-- DB gallery_db

CREATE TABLE "products"
(
    "id"          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "title"       VARCHAR(255) NOT NULL,
    "image_url"   VARCHAR(255) NOT NULL,
    "description" TEXT         NOT NULL
);
ALTER TABLE "products"
    ADD PRIMARY KEY ("id");

CREATE TABLE "stores"
(
    "id"   BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NOT NULL,
    "icon" VARCHAR(255) NOT NULL
);
ALTER TABLE "stores"
    ADD PRIMARY KEY ("id");

CREATE TABLE "links"
(
    "id"         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "product_id" BIGINT NOT NULL,
    "store_id"   BIGINT NOT NULL,
    "link"       VARCHAR(255) NOT NULL
);
ALTER TABLE "links"
    ADD PRIMARY KEY ("id");
ALTER TABLE "links"
    ADD CONSTRAINT "links_product_id_foreign" FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "links"
    ADD CONSTRAINT "links_store_id_foreign" FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
