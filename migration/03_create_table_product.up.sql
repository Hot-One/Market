CREATE TABLE "product"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(55) NOT NULL,
    "category_id" UUID REFERENCES "category"("id"),
    "barcode" VARCHAR UNIQUE NOT NULL,
    "price" NUMERIC,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);