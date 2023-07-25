CREATE TABLE "category"(
    "id" UUID NOT NULL PRIMARY KEY,
    "title" VARCHAR(50) NOT NULL,
    "parent_id" UUID REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);