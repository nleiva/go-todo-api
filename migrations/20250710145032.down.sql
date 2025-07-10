-- reverse: create index "idx_todos_deleted_at" to table: "todos"
DROP INDEX `idx_todos_deleted_at`;
-- reverse: create "todos" table
DROP TABLE `todos`;
-- reverse: create index "idx_accounts_deleted_at" to table: "accounts"
DROP INDEX `idx_accounts_deleted_at`;
-- reverse: create index "idx_accounts_email" to table: "accounts"
DROP INDEX `idx_accounts_email`;
-- reverse: create "accounts" table
DROP TABLE `accounts`;
