BEGIN;

-- User Auth
CREATE UNIQUE INDEX IF NOT EXISTS "user_email_unique_key" ON "user_authentications" ("email");
ALTER TABLE "user_authentications" ADD CONSTRAINT "user_auth_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "user_roles" ("id");

-- User Roles
CREATE UNIQUE INDEX IF NOT EXISTS "role_name_unique_key" ON "user_roles" ("name");

-- User Sessions
ALTER TABLE "user_sessions" ADD CONSTRAINT "session_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "user_authentications" ("user_id");

COMMIT;