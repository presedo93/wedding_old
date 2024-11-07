CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "profiles" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar NOT NULL,
  "picture_url" varchar,
  "completed_profile" bool NOT NULL DEFAULT false,
  "added_guests" bool NOT NULL DEFAULT false,
  "added_songs" bool NOT NULL DEFAULT false,
  "added_pictures" bool NOT NULL DEFAULT false
);

CREATE TABLE "guests" (
  "id" bigserial PRIMARY KEY,
  "profile_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "is_vegetarian" bool NOT NULL,
  "allergies" varchar[] NOT NULL,
  "needs_transport" bool NOT NULL
);

ALTER TABLE "guests" ADD CONSTRAINT "fk_profile" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE CASCADE;

CREATE TABLE "songs" (
  "id" bigserial PRIMARY KEY,
  "profile_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "name" varchar NOT NULL,
  "album" varchar NOT NULL,
  "album_picture" varchar NOT NULL
);

ALTER TABLE "songs" ADD CONSTRAINT "fk_profile" FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id") ON DELETE CASCADE;
