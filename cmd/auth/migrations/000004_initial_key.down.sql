BEGIN;

-- User Auth
DROP INDEX IF EXISTS "user_email_unique_key";
ALTER TABLE "user_authentications" DROP CONSTRAINT IF EXISTS "user_auth_role_id_fkey";

-- User Roles
DROP INDEX IF EXISTS "role_name_unique_key";

-- User Sessions
ALTER TABLE "user_sessions" DROP CONSTRAINT  IF EXISTS "session_user_id_fkey";

COMMIT;