CREATE SCHEMA IF NOT EXISTS test AUTHORIZATION postgres;
DROP TABLE IF EXISTS cat;

CREATE TABLE "cat" (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "age" INT
);

INSERT INTO cat (id, name, age) VALUES (1,'cat1',1);
INSERT INTO cat (id, name, age) VALUES (2,'cat3',2);

