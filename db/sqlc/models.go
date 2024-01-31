// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	null "gopkg.in/guregu/null.v4"
)

// Stores dishes
type Dish struct {
	ID           int64       `json:"id"`
	RestaurantID int64       `json:"restaurant_id"`
	IsAvailable  bool        `json:"is_available"`
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	Price        float64     `json:"price"`
	Cuisine      null.String `json:"cuisine"`
	ImageUrl     null.String `json:"image_url"`
}

// Stores playlist
type Playlist struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description null.String `json:"description"`
	ImageUrl    null.String `json:"image_url"`
	IsPublic    bool        `json:"is_public"`
	DeliveryDay null.String `json:"delivery_day"`
	Category    null.String `json:"category"`
	CreatedAt   time.Time   `json:"created_at"`
	AddedAt     time.Time   `json:"added_at"`
}

// Stores playlist dishes
type PlaylistDish struct {
	ID           int64     `json:"id"`
	OrderID      int64     `json:"order_id"`
	PlaylistID   int64     `json:"playlist_id"`
	DishID       int64     `json:"dish_id"`
	DishQuantity int64     `json:"dish_quantity"`
	CreatedAt    time.Time `json:"created_at"`
	AddedAt      time.Time `json:"added_at"`
}

// Stores restaurants
type Restaurant struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description null.String `json:"description"`
	Location    null.String `json:"location"`
	Cuisine     null.String `json:"cuisine"`
	ImageUrl    null.String `json:"image_url"`
}

// Stores user searches
type Search struct {
	ID      int64       `json:"id"`
	UserID  int64       `json:"user_id"`
	Keyword null.String `json:"keyword"`
}

// Stores user data
type User struct {
	ID           int64       `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
	Address      null.String `json:"address"`
}

// Stores user playlist
type UserPlaylist struct {
	ID           int64       `json:"id"`
	UserID       int64       `json:"user_id"`
	PlaylistID   int64       `json:"playlist_id"`
	DeliveryDay  null.String `json:"delivery_day"`
	DeliveryTime null.Time   `json:"delivery_time"`
	Status       null.String `json:"status"`
}
