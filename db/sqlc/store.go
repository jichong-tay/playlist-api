package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	null "gopkg.in/guregu/null.v4"
)

// Store provides all function to execute db queries and transactions
type Store interface {
	Querier
	CreatePlaylistTx(ctx context.Context, arg CreatePlaylistTxParams) (Playlist, error)
	CreatePlaylistDishTx(ctx context.Context, arg PlaylistDishTxParams) (Playlist, error)
	UpdatePlaylistDishTx(ctx context.Context, arg PlaylistDishTxParams) (Playlist, error)
}

// SQLStore provides all function to execute db queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx excutes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

/*
Transaction to create
//1. Generate a playlist based on cuisine and number of restaurant item, and budget (per restaurant item)
		- create a playlist
		- insert
//2. create a playlist, add dishes to playlist
*/

// CreatePlaylistTxParams defines the parameters to create a playlist
type CreatePlaylistTxParams struct {
	Name         string      `json:"name"`
	Description  null.String `json:"description"`
	ImageURL     null.String `json:"image_url"`
	IsPublic     bool        `json:"is_public"`
	DeliveryDay  null.String `json:"delivery_day"`
	Category     null.String `json:"category"`
	UserID       int64       `json:"user_id"`
	PlaylistID   int64       `json:"playlist_id"`
	DeliveryTime null.Time   `json:"delivery_time"`
	Status       null.String `json:"status"`
}

// CreatePlaylistTx creates a new playlist and adds it to the user's playlist
func (store *SQLStore) CreatePlaylistTx(ctx context.Context, arg CreatePlaylistTxParams) (Playlist, error) {
	var playlist Playlist
	var err error

	err = store.execTx(ctx, func(q *Queries) error {
		playlist, err = q.CreatePlaylist(ctx, CreatePlaylistParams{
			Name:        arg.Name,
			Description: arg.Description,
			ImageUrl:    arg.ImageURL,
			IsPublic:    arg.IsPublic,
			DeliveryDay: arg.DeliveryDay,
			Category:    arg.Category,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateUser_Playlist(ctx, CreateUser_PlaylistParams{
			UserID:       arg.UserID,
			PlaylistID:   playlist.ID,
			DeliveryDay:  arg.DeliveryDay,
			DeliveryTime: arg.DeliveryTime,
			Status:       arg.Status,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return playlist, err
}

// PlaylistDishTxParams defines the parameters to create a playlist with dishes
type PlaylistDishTxParams struct {
	Name         string
	Description  null.String
	ImageURL     null.String
	IsPublic     bool
	DeliveryDay  null.String
	Category     null.String
	UserID       int64
	PlaylistID   int64
	DeliveryTime null.Time
	Status       null.String
	DishItems    []RestaurantFoodItem
}

// RestaurantFoodItem defines the parameters to create a restaurant food item
type RestaurantFoodItem struct {
	RestaurantName string     `form:"restaurantName" json:"restaurantName"`
	FoodItems      []FoodItem `form:"foodItems" json:"foodItems"`
}

// FoodItem defines the parameters to create a food item
type FoodItem struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Quantity    int64   `form:"quantity" json:"quantity"`
	Price       float64 `form:"price" json:"price"`
	ImageURL    string  `form:"imageUrl" json:"imageUrl"`
	DishID      int64   `form:"dishId" json:"dishId"`
}

// CreatePlaylistDishTx creates a new playlist and adds it to the user's playlist
func (store *SQLStore) CreatePlaylistDishTx(ctx context.Context, arg PlaylistDishTxParams) (Playlist, error) {
	var playlist Playlist
	var err error

	err = store.execTx(ctx, func(q *Queries) error {
		playlist, err = q.CreatePlaylist(ctx, CreatePlaylistParams{
			Name:        arg.Name,
			Description: arg.Description,
			ImageUrl:    arg.ImageURL,
			IsPublic:    arg.IsPublic,
			DeliveryDay: arg.DeliveryDay,
			Category:    arg.Category,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateUser_Playlist(ctx, CreateUser_PlaylistParams{
			UserID:       arg.UserID,
			PlaylistID:   playlist.ID,
			DeliveryDay:  arg.DeliveryDay,
			DeliveryTime: arg.DeliveryTime,
			Status:       arg.Status,
		})
		if err != nil {
			return err
		}

		for _, restaurantFoodItem := range arg.DishItems {
			for _, foodItem := range restaurantFoodItem.FoodItems {
				_, err = q.CreatePlaylist_Dish(ctx, CreatePlaylist_DishParams{
					OrderID:      playlist.ID,
					PlaylistID:   playlist.ID,
					DishID:       foodItem.DishID,
					DishQuantity: foodItem.Quantity,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return playlist, err
}

// UpdatePlaylistDishTx updates a playlist and adds it to the user's playlist
func (store *SQLStore) UpdatePlaylistDishTx(ctx context.Context, arg PlaylistDishTxParams) (Playlist, error) {
	var playlist Playlist
	var err error

	err = store.execTx(ctx, func(q *Queries) error {
		playlist, err = q.GetPlaylist(ctx, arg.PlaylistID)
		if err != nil {
			return err
		}

		playlistDishes, err := q.DeletePlaylistDishes(ctx, arg.PlaylistID)
		if err != nil {
			return err
		}
		if len(playlistDishes) == 0 {
			return sql.ErrNoRows
		}

		_, err = q.UpdatePlaylist(ctx, UpdatePlaylistParams{
			ID:          arg.PlaylistID,
			Name:        arg.Name,
			Description: arg.Description,
			ImageUrl:    arg.ImageURL,
			IsPublic:    arg.IsPublic,
			DeliveryDay: arg.DeliveryDay,
			Category:    arg.Category,
			AddedAt:     time.Now(),
		})
		if err != nil {
			return err
		}

		_, err = q.UpdateUser_PlaylistDelivery(ctx, UpdateUser_PlaylistDeliveryParams{
			UserID:       arg.UserID,
			PlaylistID:   arg.PlaylistID,
			DeliveryDay:  arg.DeliveryDay,
			DeliveryTime: arg.DeliveryTime,
			Status:       arg.Status,
		})
		if err != nil {
			return err
		}

		for _, restaurantFoodItem := range arg.DishItems {
			for _, foodItem := range restaurantFoodItem.FoodItems {
				_, err := q.CreatePlaylist_Dish(ctx, CreatePlaylist_DishParams{
					OrderID:      playlist.ID,
					PlaylistID:   playlist.ID,
					DishID:       foodItem.DishID,
					DishQuantity: foodItem.Quantity,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return playlist, err
}
