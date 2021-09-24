BEGIN;

CREATE TABLE IF NOT EXISTS "products" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
   
    "name" text NOT NULL,
    "description" text NOT NULL,
    "price" bigint NOT NULL,
    "qty" bigint NOT NULL DEFAULT 0,

    "created_at" bigint NOT NULL,
    "created_by" UUID NOT NULL,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

COMMIT;