package provider

import (
	"github.com/akosgarai/go_example_service/database"
	"gopkg.in/gin-gonic/gin.v1"
)

type WatchObject struct {
	Category int    `form:"category" json:"category" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
	Price    int    `form:"price" json:"price" binding:"required"`
}
type ServiceObject struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	Invoice     string `form:"invoice" json:"invoice" binding:"required"`
	Bankaccount string `form:"bankaccount" json:"bankaccount" binding:"required"`
}

func isCategoryValid(id int) bool {
	if id > 10 && id < 20 {
		return true
	}
	return false
}
func isTypeValid(t string) bool {
	if t == "s" {
		return true
	}
	return false
}
func isPriceValid(price int) bool {
	if price > 2000 && price < 500000 {
		return true
	}
	return false
}
func ValidateParams(p WatchObject) gin.H {
	if isCategoryValid(p.Category) && isTypeValid(p.Type) && isPriceValid(p.Price) {
		return gin.H{
			"params":   "name,address,invoice,bankaccount",
			"provider": "provider company",
			"price":    1900,
		}
	}
	return nil
}

func ValidateStoreParams(s database.StoreObject) bool {
	return true
}
