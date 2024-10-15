CREATE TABLE "guests" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar(12) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "is_vegetarian" bool NOT NULL,
  "allergies" varchar[] NOT NULL,
  "needs_transport" bool NOT NULL
);

CREATE INDEX ON "guests" ("user_id");

CREATE TABLE "songs" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar(12) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar NOT NULL,
  "album" varchar NOT NULL,
  "album_picture" varchar NOT NULL
);

CREATE INDEX ON "songs" ("user_id");
