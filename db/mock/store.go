// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jichong-tay/foodpanda-playlist-api/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go github.com/jichong-tay/foodpanda-playlist-api/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
	gomock "go.uber.org/mock/gomock"
	null "gopkg.in/guregu/null.v4"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateDish mocks base method.
func (m *MockStore) CreateDish(arg0 context.Context, arg1 db.CreateDishParams) (db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDish", arg0, arg1)
	ret0, _ := ret[0].(db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDish indicates an expected call of CreateDish.
func (mr *MockStoreMockRecorder) CreateDish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDish", reflect.TypeOf((*MockStore)(nil).CreateDish), arg0, arg1)
}

// CreatePlaylist mocks base method.
func (m *MockStore) CreatePlaylist(arg0 context.Context, arg1 db.CreatePlaylistParams) (db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", arg0, arg1)
	ret0, _ := ret[0].(db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockStoreMockRecorder) CreatePlaylist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockStore)(nil).CreatePlaylist), arg0, arg1)
}

// CreatePlaylistTx mocks base method.
func (m *MockStore) CreatePlaylistTx(arg0 context.Context, arg1 db.CreatePlaylistTxParams) (db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylistTx", arg0, arg1)
	ret0, _ := ret[0].(db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylistTx indicates an expected call of CreatePlaylistTx.
func (mr *MockStoreMockRecorder) CreatePlaylistTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylistTx", reflect.TypeOf((*MockStore)(nil).CreatePlaylistTx), arg0, arg1)
}

// CreatePlaylist_Dish mocks base method.
func (m *MockStore) CreatePlaylist_Dish(arg0 context.Context, arg1 db.CreatePlaylist_DishParams) (db.PlaylistDish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist_Dish", arg0, arg1)
	ret0, _ := ret[0].(db.PlaylistDish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist_Dish indicates an expected call of CreatePlaylist_Dish.
func (mr *MockStoreMockRecorder) CreatePlaylist_Dish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist_Dish", reflect.TypeOf((*MockStore)(nil).CreatePlaylist_Dish), arg0, arg1)
}

// CreateRestaurant mocks base method.
func (m *MockStore) CreateRestaurant(arg0 context.Context, arg1 db.CreateRestaurantParams) (db.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestaurant", arg0, arg1)
	ret0, _ := ret[0].(db.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRestaurant indicates an expected call of CreateRestaurant.
func (mr *MockStoreMockRecorder) CreateRestaurant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurant", reflect.TypeOf((*MockStore)(nil).CreateRestaurant), arg0, arg1)
}

// CreateSearch mocks base method.
func (m *MockStore) CreateSearch(arg0 context.Context, arg1 db.CreateSearchParams) (db.Search, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearch", arg0, arg1)
	ret0, _ := ret[0].(db.Search)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSearch indicates an expected call of CreateSearch.
func (mr *MockStoreMockRecorder) CreateSearch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearch", reflect.TypeOf((*MockStore)(nil).CreateSearch), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUser_Playlist mocks base method.
func (m *MockStore) CreateUser_Playlist(arg0 context.Context, arg1 db.CreateUser_PlaylistParams) (db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser_Playlist", arg0, arg1)
	ret0, _ := ret[0].(db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser_Playlist indicates an expected call of CreateUser_Playlist.
func (mr *MockStoreMockRecorder) CreateUser_Playlist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser_Playlist", reflect.TypeOf((*MockStore)(nil).CreateUser_Playlist), arg0, arg1)
}

// DeleteDish mocks base method.
func (m *MockStore) DeleteDish(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDish", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDish indicates an expected call of DeleteDish.
func (mr *MockStoreMockRecorder) DeleteDish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDish", reflect.TypeOf((*MockStore)(nil).DeleteDish), arg0, arg1)
}

// DeletePlaylist mocks base method.
func (m *MockStore) DeletePlaylist(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylist indicates an expected call of DeletePlaylist.
func (mr *MockStoreMockRecorder) DeletePlaylist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist", reflect.TypeOf((*MockStore)(nil).DeletePlaylist), arg0, arg1)
}

// DeletePlaylist_Dish mocks base method.
func (m *MockStore) DeletePlaylist_Dish(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist_Dish", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylist_Dish indicates an expected call of DeletePlaylist_Dish.
func (mr *MockStoreMockRecorder) DeletePlaylist_Dish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist_Dish", reflect.TypeOf((*MockStore)(nil).DeletePlaylist_Dish), arg0, arg1)
}

// DeleteRestaurant mocks base method.
func (m *MockStore) DeleteRestaurant(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRestaurant", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRestaurant indicates an expected call of DeleteRestaurant.
func (mr *MockStoreMockRecorder) DeleteRestaurant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRestaurant", reflect.TypeOf((*MockStore)(nil).DeleteRestaurant), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUser_Playlist mocks base method.
func (m *MockStore) DeleteUser_Playlist(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser_Playlist", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser_Playlist indicates an expected call of DeleteUser_Playlist.
func (mr *MockStoreMockRecorder) DeleteUser_Playlist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser_Playlist", reflect.TypeOf((*MockStore)(nil).DeleteUser_Playlist), arg0, arg1)
}

// GetDish mocks base method.
func (m *MockStore) GetDish(arg0 context.Context, arg1 int64) (db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDish", arg0, arg1)
	ret0, _ := ret[0].(db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDish indicates an expected call of GetDish.
func (mr *MockStoreMockRecorder) GetDish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDish", reflect.TypeOf((*MockStore)(nil).GetDish), arg0, arg1)
}

// GetPlaylist mocks base method.
func (m *MockStore) GetPlaylist(arg0 context.Context, arg1 int64) (db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist", arg0, arg1)
	ret0, _ := ret[0].(db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylist indicates an expected call of GetPlaylist.
func (mr *MockStoreMockRecorder) GetPlaylist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist", reflect.TypeOf((*MockStore)(nil).GetPlaylist), arg0, arg1)
}

// GetPlaylist_Dish mocks base method.
func (m *MockStore) GetPlaylist_Dish(arg0 context.Context, arg1 int64) (db.PlaylistDish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist_Dish", arg0, arg1)
	ret0, _ := ret[0].(db.PlaylistDish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylist_Dish indicates an expected call of GetPlaylist_Dish.
func (mr *MockStoreMockRecorder) GetPlaylist_Dish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist_Dish", reflect.TypeOf((*MockStore)(nil).GetPlaylist_Dish), arg0, arg1)
}

// GetRestaurant mocks base method.
func (m *MockStore) GetRestaurant(arg0 context.Context, arg1 int64) (db.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", arg0, arg1)
	ret0, _ := ret[0].(db.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockStoreMockRecorder) GetRestaurant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockStore)(nil).GetRestaurant), arg0, arg1)
}

// GetSearch mocks base method.
func (m *MockStore) GetSearch(arg0 context.Context, arg1 int64) (db.Search, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearch", arg0, arg1)
	ret0, _ := ret[0].(db.Search)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSearch indicates an expected call of GetSearch.
func (mr *MockStoreMockRecorder) GetSearch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearch", reflect.TypeOf((*MockStore)(nil).GetSearch), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserPlaylistByPlaylistID mocks base method.
func (m *MockStore) GetUserPlaylistByPlaylistID(arg0 context.Context, arg1 int64) (db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPlaylistByPlaylistID", arg0, arg1)
	ret0, _ := ret[0].(db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPlaylistByPlaylistID indicates an expected call of GetUserPlaylistByPlaylistID.
func (mr *MockStoreMockRecorder) GetUserPlaylistByPlaylistID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPlaylistByPlaylistID", reflect.TypeOf((*MockStore)(nil).GetUserPlaylistByPlaylistID), arg0, arg1)
}

// GetUser_Playlist mocks base method.
func (m *MockStore) GetUser_Playlist(arg0 context.Context, arg1 int64) (db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser_Playlist", arg0, arg1)
	ret0, _ := ret[0].(db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser_Playlist indicates an expected call of GetUser_Playlist.
func (mr *MockStoreMockRecorder) GetUser_Playlist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser_Playlist", reflect.TypeOf((*MockStore)(nil).GetUser_Playlist), arg0, arg1)
}

// ListDishes mocks base method.
func (m *MockStore) ListDishes(arg0 context.Context, arg1 db.ListDishesParams) ([]db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDishes", arg0, arg1)
	ret0, _ := ret[0].([]db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDishes indicates an expected call of ListDishes.
func (mr *MockStoreMockRecorder) ListDishes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDishes", reflect.TypeOf((*MockStore)(nil).ListDishes), arg0, arg1)
}

// ListPlaylistByCategory mocks base method.
func (m *MockStore) ListPlaylistByCategory(arg0 context.Context, arg1 string) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylistByCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylistByCategory indicates an expected call of ListPlaylistByCategory.
func (mr *MockStoreMockRecorder) ListPlaylistByCategory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylistByCategory", reflect.TypeOf((*MockStore)(nil).ListPlaylistByCategory), arg0, arg1)
}

// ListPlaylistPublicAndCategory mocks base method.
func (m *MockStore) ListPlaylistPublicAndCategory(arg0 context.Context, arg1 db.ListPlaylistPublicAndCategoryParams) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylistPublicAndCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylistPublicAndCategory indicates an expected call of ListPlaylistPublicAndCategory.
func (mr *MockStoreMockRecorder) ListPlaylistPublicAndCategory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylistPublicAndCategory", reflect.TypeOf((*MockStore)(nil).ListPlaylistPublicAndCategory), arg0, arg1)
}

// ListPlaylistPublicAndCategoryAll mocks base method.
func (m *MockStore) ListPlaylistPublicAndCategoryAll(arg0 context.Context) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylistPublicAndCategoryAll", arg0)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylistPublicAndCategoryAll indicates an expected call of ListPlaylistPublicAndCategoryAll.
func (mr *MockStoreMockRecorder) ListPlaylistPublicAndCategoryAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylistPublicAndCategoryAll", reflect.TypeOf((*MockStore)(nil).ListPlaylistPublicAndCategoryAll), arg0)
}

// ListPlaylist_Dishes mocks base method.
func (m *MockStore) ListPlaylist_Dishes(arg0 context.Context, arg1 db.ListPlaylist_DishesParams) ([]db.PlaylistDish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylist_Dishes", arg0, arg1)
	ret0, _ := ret[0].([]db.PlaylistDish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylist_Dishes indicates an expected call of ListPlaylist_Dishes.
func (mr *MockStoreMockRecorder) ListPlaylist_Dishes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylist_Dishes", reflect.TypeOf((*MockStore)(nil).ListPlaylist_Dishes), arg0, arg1)
}

// ListPlaylist_DishesByPlaylistID mocks base method.
func (m *MockStore) ListPlaylist_DishesByPlaylistID(arg0 context.Context, arg1 int64) ([]db.PlaylistDish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylist_DishesByPlaylistID", arg0, arg1)
	ret0, _ := ret[0].([]db.PlaylistDish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylist_DishesByPlaylistID indicates an expected call of ListPlaylist_DishesByPlaylistID.
func (mr *MockStoreMockRecorder) ListPlaylist_DishesByPlaylistID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylist_DishesByPlaylistID", reflect.TypeOf((*MockStore)(nil).ListPlaylist_DishesByPlaylistID), arg0, arg1)
}

// ListPlaylists mocks base method.
func (m *MockStore) ListPlaylists(arg0 context.Context, arg1 db.ListPlaylistsParams) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylists", arg0, arg1)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylists indicates an expected call of ListPlaylists.
func (mr *MockStoreMockRecorder) ListPlaylists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylists", reflect.TypeOf((*MockStore)(nil).ListPlaylists), arg0, arg1)
}

// ListPlaylistsByUserID mocks base method.
func (m *MockStore) ListPlaylistsByUserID(arg0 context.Context, arg1 db.ListPlaylistsByUserIDParams) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylistsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylistsByUserID indicates an expected call of ListPlaylistsByUserID.
func (mr *MockStoreMockRecorder) ListPlaylistsByUserID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylistsByUserID", reflect.TypeOf((*MockStore)(nil).ListPlaylistsByUserID), arg0, arg1)
}

// ListPlaylistsByUserIDAll mocks base method.
func (m *MockStore) ListPlaylistsByUserIDAll(arg0 context.Context, arg1 int64) ([]db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylistsByUserIDAll", arg0, arg1)
	ret0, _ := ret[0].([]db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylistsByUserIDAll indicates an expected call of ListPlaylistsByUserIDAll.
func (mr *MockStoreMockRecorder) ListPlaylistsByUserIDAll(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylistsByUserIDAll", reflect.TypeOf((*MockStore)(nil).ListPlaylistsByUserIDAll), arg0, arg1)
}

// ListRestaurantNameByDishID mocks base method.
func (m *MockStore) ListRestaurantNameByDishID(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRestaurantNameByDishID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestaurantNameByDishID indicates an expected call of ListRestaurantNameByDishID.
func (mr *MockStoreMockRecorder) ListRestaurantNameByDishID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurantNameByDishID", reflect.TypeOf((*MockStore)(nil).ListRestaurantNameByDishID), arg0, arg1)
}

// ListRestaurants mocks base method.
func (m *MockStore) ListRestaurants(arg0 context.Context, arg1 db.ListRestaurantsParams) ([]db.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRestaurants", arg0, arg1)
	ret0, _ := ret[0].([]db.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestaurants indicates an expected call of ListRestaurants.
func (mr *MockStoreMockRecorder) ListRestaurants(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurants", reflect.TypeOf((*MockStore)(nil).ListRestaurants), arg0, arg1)
}

// ListSearches mocks base method.
func (m *MockStore) ListSearches(arg0 context.Context, arg1 db.ListSearchesParams) ([]db.Search, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSearches", arg0, arg1)
	ret0, _ := ret[0].([]db.Search)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearches indicates an expected call of ListSearches.
func (mr *MockStoreMockRecorder) ListSearches(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearches", reflect.TypeOf((*MockStore)(nil).ListSearches), arg0, arg1)
}

// ListStatusByPlaylistID mocks base method.
func (m *MockStore) ListStatusByPlaylistID(arg0 context.Context, arg1 int64) (null.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatusByPlaylistID", arg0, arg1)
	ret0, _ := ret[0].(null.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStatusByPlaylistID indicates an expected call of ListStatusByPlaylistID.
func (mr *MockStoreMockRecorder) ListStatusByPlaylistID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatusByPlaylistID", reflect.TypeOf((*MockStore)(nil).ListStatusByPlaylistID), arg0, arg1)
}

// ListUser_Playlists mocks base method.
func (m *MockStore) ListUser_Playlists(arg0 context.Context, arg1 db.ListUser_PlaylistsParams) ([]db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser_Playlists", arg0, arg1)
	ret0, _ := ret[0].([]db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser_Playlists indicates an expected call of ListUser_Playlists.
func (mr *MockStoreMockRecorder) ListUser_Playlists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser_Playlists", reflect.TypeOf((*MockStore)(nil).ListUser_Playlists), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// UpdateDish mocks base method.
func (m *MockStore) UpdateDish(arg0 context.Context, arg1 db.UpdateDishParams) (db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDish", arg0, arg1)
	ret0, _ := ret[0].(db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDish indicates an expected call of UpdateDish.
func (mr *MockStoreMockRecorder) UpdateDish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDish", reflect.TypeOf((*MockStore)(nil).UpdateDish), arg0, arg1)
}

// UpdatePlaylist mocks base method.
func (m *MockStore) UpdatePlaylist(arg0 context.Context, arg1 db.UpdatePlaylistParams) (db.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylist", arg0, arg1)
	ret0, _ := ret[0].(db.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePlaylist indicates an expected call of UpdatePlaylist.
func (mr *MockStoreMockRecorder) UpdatePlaylist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylist", reflect.TypeOf((*MockStore)(nil).UpdatePlaylist), arg0, arg1)
}

// UpdatePlaylist_Dish mocks base method.
func (m *MockStore) UpdatePlaylist_Dish(arg0 context.Context, arg1 db.UpdatePlaylist_DishParams) (db.PlaylistDish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylist_Dish", arg0, arg1)
	ret0, _ := ret[0].(db.PlaylistDish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePlaylist_Dish indicates an expected call of UpdatePlaylist_Dish.
func (mr *MockStoreMockRecorder) UpdatePlaylist_Dish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylist_Dish", reflect.TypeOf((*MockStore)(nil).UpdatePlaylist_Dish), arg0, arg1)
}

// UpdateRestaurant mocks base method.
func (m *MockStore) UpdateRestaurant(arg0 context.Context, arg1 db.UpdateRestaurantParams) (db.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRestaurant", arg0, arg1)
	ret0, _ := ret[0].(db.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRestaurant indicates an expected call of UpdateRestaurant.
func (mr *MockStoreMockRecorder) UpdateRestaurant(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRestaurant", reflect.TypeOf((*MockStore)(nil).UpdateRestaurant), arg0, arg1)
}

// UpdateStatusForUser_Playlist mocks base method.
func (m *MockStore) UpdateStatusForUser_Playlist(arg0 context.Context, arg1 db.UpdateStatusForUser_PlaylistParams) ([]db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusForUser_Playlist", arg0, arg1)
	ret0, _ := ret[0].([]db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatusForUser_Playlist indicates an expected call of UpdateStatusForUser_Playlist.
func (mr *MockStoreMockRecorder) UpdateStatusForUser_Playlist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusForUser_Playlist", reflect.TypeOf((*MockStore)(nil).UpdateStatusForUser_Playlist), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUser_Playlist mocks base method.
func (m *MockStore) UpdateUser_Playlist(arg0 context.Context, arg1 db.UpdateUser_PlaylistParams) (db.UserPlaylist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser_Playlist", arg0, arg1)
	ret0, _ := ret[0].(db.UserPlaylist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser_Playlist indicates an expected call of UpdateUser_Playlist.
func (mr *MockStoreMockRecorder) UpdateUser_Playlist(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser_Playlist", reflect.TypeOf((*MockStore)(nil).UpdateUser_Playlist), arg0, arg1)
}
