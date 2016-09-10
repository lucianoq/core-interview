DROP TABLE IF EXISTS encrypted;

CREATE TABLE `encrypted` (
  `id`         VARCHAR(30) PRIMARY KEY,
  `ciphertext` VARCHAR(1000) NOT NULL
);
