CREATE TYPE "role" AS ENUM (
  'user',
  'admin'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "role" role DEFAULT 'user',
  "companions" int DEFAULT 1
);

CREATE TABLE "guests" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "name" varchar NOT NULL,
  "user_id" int,
  "is_vegetarian" bool,
  "allergies" varchar[],
  "is_using_bus" bool
);

CREATE TABLE "songs" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "name" varchar NOT NULL,
  "album" varchar,
  "album_picture" varchar
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
