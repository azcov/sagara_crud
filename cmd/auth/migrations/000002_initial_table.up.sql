BEGIN;

CREATE TABLE IF NOT EXISTS "user_roles" (
    "id" serial NOT NULL PRIMARY KEY,

    "name" text NOT NULL,

    "created_at" bigint,
    "updated_at" bigint
);

CREATE TABLE IF NOT EXISTS "user_authentications" (
    "user_id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    "role_id" int NOT NULL,
    "name" text NOT NULL,
    "email" bytea NOT NULL, -- must encrypted
    "password" text NOT NULL, -- must encrypted

    "created_at" bigint,
    "updated_at" bigint,
    "deleted_at" bigint
);
CREATE TABLE IF NOT EXISTS "user_sessions" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,

    "last_login" bigint NOT NULL,
    "access_token" text NOT NULL, -- encrypted
    "refresh_token" text NOT NULL, -- encrypted

    "is_active" boolean NOT NULL DEFAULT false,

    "last_ip_address" text NOT NULL,

    "updated_at" bigint,
    "deleted_at" bigint
);

COMMIT;