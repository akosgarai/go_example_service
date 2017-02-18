package database

import (
	"encoding/json"
	"errors"
	"github.com/akosgarai/go_akos_httpd/store"
	"strconv"
	"time"
)

type StoreObject struct {
	Category    int    `form:"category" json:"category" binding:"required"`
	Type        string `form:"type" json:"type" binding:"required"`
	Price       int    `form:"price" json:"price" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	Invoice     string `form:"invoice" json:"invoice" binding:"required"`
	Bankaccount string `form:"bankaccount" json:"bankaccount" binding:"required"`
	DeliveryID  int64  `form:"deliveryid" json:"deliveryid"`
}

type DataBase struct {
	Store store.Store
}

// New returns a new Database.
func New() *DataBase {
	store := store.New()
	store.Open()
	return &DataBase{
		Store: *store,
	}
}

func (db *DataBase) StoreData(s StoreObject) (int64, error) {
	if s.DeliveryID > 0 {
		return 0, errors.New("Not new Delivery")
	}
	storeId := db.CreateServiceId()
	s.DeliveryID = storeId
	jsonString, err := json.Marshal(s)
	if err != nil {
		return storeId, err
	}
	return storeId, db.Store.Set(FormatKey(storeId), string(jsonString[:]))
}
func (db *DataBase) UpdateData(s StoreObject) (int64, error) {
	storeId := s.DeliveryID
	jsonString, err := json.Marshal(s)
	if err != nil {
		return storeId, err
	}
	return storeId, db.Store.Set(FormatKey(storeId), string(jsonString[:]))
}
func (db *DataBase) GetData(key string) (string, error) {
	return db.Store.Get(key)
}

func (db *DataBase) CreateServiceId() int64 {
	return time.Now().Unix()
}

func FormatKey(key int64) string {
	return strconv.FormatInt(key, 10)
}
