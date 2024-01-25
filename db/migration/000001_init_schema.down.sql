-- Drop foreign keys before dropping tables
ALTER TABLE "playlist_dishes" DROP CONSTRAINT IF EXISTS "playlist_dishes_dish_id_fkey";
ALTER TABLE "playlist_dishes" DROP CONSTRAINT IF EXISTS "playlist_dishes_playlist_id_fkey";
ALTER TABLE "user_playlists" DROP CONSTRAINT IF EXISTS "user_playlists_playlist_id_fkey";
ALTER TABLE "user_playlists" DROP CONSTRAINT IF EXISTS "user_playlists_user_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "playlist_dishes";
DROP TABLE IF EXISTS "dishes";
DROP TABLE IF EXISTS "restaurants";
DROP TABLE IF EXISTS "playlists";
DROP TABLE IF EXISTS "user_playlists";
DROP TABLE IF EXISTS "searches";
DROP TABLE IF EXISTS "users";