// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateDish(ctx context.Context, arg CreateDishParams) (Dish, error)
	CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error)
	CreatePlaylist_Dish(ctx context.Context, arg CreatePlaylist_DishParams) (PlaylistDish, error)
	CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (Restaurant, error)
	CreateSearch(ctx context.Context, arg CreateSearchParams) (Search, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUser_Playlist(ctx context.Context, arg CreateUser_PlaylistParams) (UserPlaylist, error)
	DeleteDish(ctx context.Context, id int64) error
	DeletePlaylist(ctx context.Context, id int64) error
	DeletePlaylist_Dish(ctx context.Context, id int64) error
	DeleteRestaurant(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteUser_Playlist(ctx context.Context, id int64) error
	GetDish(ctx context.Context, id int64) (Dish, error)
	GetPlaylist(ctx context.Context, id int64) (Playlist, error)
	GetPlaylist_Dish(ctx context.Context, id int64) (PlaylistDish, error)
	GetRestaurant(ctx context.Context, id int64) (Restaurant, error)
	GetSearch(ctx context.Context, id int64) (Search, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUser_Playlist(ctx context.Context, id int64) (UserPlaylist, error)
	ListDishes(ctx context.Context, arg ListDishesParams) ([]Dish, error)
	ListPlaylistByCategory(ctx context.Context, lower string) ([]Playlist, error)
	ListPlaylistPublicAndCategory(ctx context.Context, arg ListPlaylistPublicAndCategoryParams) ([]Playlist, error)
	ListPlaylistPublicAndCategoryAll(ctx context.Context) ([]Playlist, error)
	ListPlaylist_Dishes(ctx context.Context, arg ListPlaylist_DishesParams) ([]PlaylistDish, error)
	ListPlaylist_DishesByPlaylistID(ctx context.Context, playlistID int64) ([]PlaylistDish, error)
	ListPlaylists(ctx context.Context, arg ListPlaylistsParams) ([]Playlist, error)
	ListPlaylistsByUserID(ctx context.Context, arg ListPlaylistsByUserIDParams) ([]Playlist, error)
	ListPlaylistsByUserIDAll(ctx context.Context, userID int64) ([]Playlist, error)
	ListPublicPlaylist(ctx context.Context) ([]Playlist, error)
	ListRestaurantNameByDishID(ctx context.Context, id int64) (string, error)
	ListRestaurants(ctx context.Context, arg ListRestaurantsParams) ([]Restaurant, error)
	ListSearches(ctx context.Context, arg ListSearchesParams) ([]Search, error)
	ListUser_Playlists(ctx context.Context, arg ListUser_PlaylistsParams) ([]UserPlaylist, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateDish(ctx context.Context, arg UpdateDishParams) (Dish, error)
	UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) (Playlist, error)
	UpdatePlaylist_Dish(ctx context.Context, arg UpdatePlaylist_DishParams) (PlaylistDish, error)
	UpdateRestaurant(ctx context.Context, arg UpdateRestaurantParams) (Restaurant, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUser_Playlist(ctx context.Context, arg UpdateUser_PlaylistParams) (UserPlaylist, error)
}

var _ Querier = (*Queries)(nil)
