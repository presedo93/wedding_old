CREATE TYPE "role" AS ENUM (
  'user',
  'admin'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "role" role NOT NULL DEFAULT 'user',
  "companions" bigint DEFAULT 1 NOT NULL
);

CREATE TABLE "guests" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "is_vegetarian" bool NOT NULL,
  "allergies" varchar[] NOT NULL,
  "is_using_bus" bool NOT NULL
);

CREATE TABLE "songs" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar NOT NULL,
  "album" varchar NOT NULL,
  "album_picture" varchar NOT NULL
);

CREATE TABLE "user_songs" (
  "user_id" bigserial,
  "song_id" bigserial,
  PRIMARY KEY ("user_id", "song_id")
);

CREATE INDEX ON "guests" ("user_id");

ALTER TABLE "guests" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_songs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_songs" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id");
