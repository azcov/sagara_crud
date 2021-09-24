BEGIN;

CREATE TABLE IF NOT EXISTS "articles" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    "title" text NOT NULL,
    "description" text NOT NULL,

    "created_at" bigint NOT NULL,
    "created_by" UUID NOT NULL,
    "updated_at" bigint,
    "updated_by" UUID,
    "deleted_at" bigint,
    "deleted_by" UUID
);

-- CREATE TABLE IF NOT EXISTS "comments" (
--     "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
--     "article_id" UUID NOT NULL,
--     "description" text NOT NULL,

--     "created_at" bigint NOT NULL,
--     "created_by" UUID NOT NULL,
--     "updated_at" bigint,
--     "updated_by" UUID,
--     "deleted_at" bigint,
--     "deleted_by" UUID
-- );




COMMIT;