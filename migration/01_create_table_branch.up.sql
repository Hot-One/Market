CREATE TABLE "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(45) NOT NULL,
    "address" VARCHAR(55),
    "phone_number" VARCHAR(55),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);



