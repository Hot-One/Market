CREATE TABLE "storage_coming"(
    "id" UUID NOT NULL PRIMARY KEY,
    "coming_id" VARCHAR NOT NULL,
    "branch_id" UUID REFERENCES branch(id),
    "status" VARCHAR DEFAULT 'in process',
    "date_time" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPstorage_income
);

CREATE TABLE "income_products"(
    "id" UUID NOT NULL PRIMARY KEY,
    "category_id" UUID REFERENCES "category"("id"),
    "name" VARCHAR NOT NULL,
    "quantity" NUMERIC,
    "price" NUMERIC,
    "total_price" NUMERIC,
    "storage_coming_id" UUID REFERENCES storage_coming("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);