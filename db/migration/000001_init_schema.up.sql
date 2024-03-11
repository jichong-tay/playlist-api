CREATE TABLE "searches" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "keyword" varchar
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "address" varchar,
  "uuid" varchar NOT NULL

);

CREATE TABLE "user_playlists" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "playlist_id" bigserial NOT NULL,
  "delivery_day" varchar,
  "delivery_time" time,
  "status" varchar
);

CREATE TABLE "playlists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "image_url" varchar,
  "is_public" bool NOT NULL,
  "delivery_day" varchar,
  "category" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "added_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "playlist_dishes" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigserial NOT NULL,
  "playlist_id" bigserial NOT NULL,
  "dish_id" bigserial NOT NULL,
  "dish_quantity" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "added_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "dishes" (
  "id" bigserial PRIMARY KEY,
  "restaurant_id" bigserial NOT NULL,
  "is_available" bool NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "price" decimal(10, 2) NOT NULL,
  "cuisine" varchar,
  "image_url" varchar
);

CREATE TABLE "restaurants" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "location" varchar,
  "cuisine" varchar,
  "image_url" varchar
);

COMMENT ON TABLE "searches" IS 'Stores user searches';

COMMENT ON TABLE "users" IS 'Stores user data';

COMMENT ON TABLE "user_playlists" IS 'Stores user playlist';

COMMENT ON TABLE "playlists" IS 'Stores playlist';

COMMENT ON TABLE "playlist_dishes" IS 'Stores playlist dishes';

COMMENT ON TABLE "dishes" IS 'Stores dishes';

COMMENT ON TABLE "restaurants" IS 'Stores restaurants';

ALTER TABLE "searches" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_playlists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_playlists" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");

ALTER TABLE "playlist_dishes" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");

ALTER TABLE "playlist_dishes" ADD FOREIGN KEY ("dish_id") REFERENCES "dishes" ("id");

ALTER TABLE "dishes" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurants" ("id");
