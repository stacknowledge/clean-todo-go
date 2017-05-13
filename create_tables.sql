CREATE TABLE `users` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `email` VARCHAR(64) UNIQUE,
    `password` VARCHAR(64) NOT NULL
);

CREATE TABLE `todos` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(120) UNIQUE,
    `expires_at` VARCHAR(20) NOT NULL,
    `user_id` INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
