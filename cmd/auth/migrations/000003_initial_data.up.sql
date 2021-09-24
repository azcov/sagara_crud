BEGIN;

INSERT INTO "user_roles" (id, name, created_at, updated_at) VALUES
(1, 'admin', extract(epoch from now()), extract(epoch from now())),
(2, 'member', extract(epoch from now()), extract(epoch from now()));


COMMIT;