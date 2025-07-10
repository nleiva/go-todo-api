-- create "accounts" table
CREATE TABLE `accounts` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `firstname` text NULL,
  `lastname` text NULL,
  `token_secret` varchar NULL,
  `permission` integer NULL DEFAULT 0
);
-- create index "idx_accounts_email" to table: "accounts"
CREATE UNIQUE INDEX `idx_accounts_email` ON `accounts` (`email`);
-- create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX `idx_accounts_deleted_at` ON `accounts` (`deleted_at`);
-- create "todos" table
CREATE TABLE `todos` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `title` text NOT NULL,
  `description` text NULL,
  `completed` numeric NULL DEFAULT false,
  `completed_at` datetime NULL,
  `account_id` integer NOT NULL,
  CONSTRAINT `fk_accounts_todos` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_todos_deleted_at" to table: "todos"
CREATE INDEX `idx_todos_deleted_at` ON `todos` (`deleted_at`);
