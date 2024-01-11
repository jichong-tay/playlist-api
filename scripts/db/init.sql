CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "password_hash" varchar,
  "address" varchar
);

CREATE TABLE "user_playlists" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "playlist_id" int,
  "delivery_day" varchar,
  "delivery_time" varchar,
  "status" varchar
);

CREATE TABLE "restaurants" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "location" varchar,
  "cuisine" varchar,
  "image_url" varchar
);

CREATE TABLE "restaurant_items" (
  "id" bigserial PRIMARY KEY,
  "restaurant_id" int,
  "is_available" bool,
  "name" varchar,
  "description" varchar,
  "price" decimal(10, 2),
  "cuisine" varchar,
  "image_url" varchar
);

CREATE TABLE "playlists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "image_url" varchar,
  "is_public" bool,
  "delivery_day" varchar,
  "category" varchar
);

CREATE TABLE "playlist_restaurant_items" (
  "id" bigserial PRIMARY KEY,
  "playlist_id" int,
  "restaurant_item_id" int,
  "restaurant_item_quantity" int,
  "added_at" timestamptz DEFAULT (now())
);

COMMENT ON TABLE "users" IS 'Stores user data';

COMMENT ON TABLE "user_playlists" IS 'Stores user playlist';

COMMENT ON TABLE "restaurants" IS 'Stores restaurants';

COMMENT ON TABLE "restaurant_items" IS 'Stores restaurant items';

COMMENT ON TABLE "playlists" IS 'Stores playlist';

COMMENT ON TABLE "playlist_restaurant_items" IS 'Stores playlist restaurant items';

ALTER TABLE "user_playlists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_playlists" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");

ALTER TABLE "restaurant_items" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurants" ("id");

ALTER TABLE "playlist_restaurant_items" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");

ALTER TABLE "playlist_restaurant_items" ADD FOREIGN KEY ("restaurant_item_id") REFERENCES "restaurant_items" ("id");
