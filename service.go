package main

import (
	"github.com/akosgarai/go_example_service/database"
	"github.com/akosgarai/go_example_service/provider"
	"gopkg.in/gin-gonic/gin.v1"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"watch": "category,type,price",
	})
}

var db = database.New()

func checkHandler(c *gin.Context) {
	var json provider.WatchObject
	if c.BindJSON(&json) == nil {
		if resp := provider.ValidateParams(json); resp != nil {
			c.JSON(200, resp)
			return
		} else {
			c.JSON(500, gin.H{
				"error": "Validation error",
			})
			return
		}
	} else {
		c.JSON(500, gin.H{
			"error": "checkHandler bind to watch object",
		})
		return
	}
}
func createHandler(c *gin.Context) {
	var dbJson database.StoreObject
	if c.BindJSON(&dbJson) != nil {
		c.JSON(500, gin.H{
			"error": "checkHandler bind to store object",
		})
		return
	}
	if provider.ValidateStoreParams(dbJson) {
		/*
		   During development period the database is mocked with this stuff.
		*/
		storeId, err := db.StoreData(dbJson)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"storeId": storeId,
		})
		return
	} else {
		c.JSON(500, gin.H{
			"error": "Validation error",
		})
		return
	}
}
func updateHandler(c *gin.Context) {
	var dbJson database.StoreObject
	if c.BindJSON(&dbJson) != nil {
		c.JSON(500, gin.H{
			"error": "checkHandler bind to store object",
		})
		return
	}
	if provider.ValidateStoreParams(dbJson) {
		id := database.FormatKey(dbJson.DeliveryID)
		val, err := db.GetData(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		if val == "" {
			c.JSON(500, gin.H{
				"error": "Record not found",
			})
			return
		} else {
			storeId, err := db.UpdateData(dbJson)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"success": "update",
				"storeid": storeId,
			})
			return
		}
	} else {
		c.JSON(500, gin.H{
			"error": "Validation error",
		})
		return
	}
}
func deleteHandler(c *gin.Context) {
}
func statusHandler(c *gin.Context) {
	id := c.Param("storeid")
	val, err := db.GetData(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if val == "" {
		c.JSON(500, gin.H{
			"error": "Record not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"storeId":   id,
		"soreValue": val,
	})
	return
}
func main() {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	router.POST("/check", checkHandler)
	router.POST("/create", createHandler)
	router.POST("/update", updateHandler)
	router.POST("/delete", deleteHandler)
	router.GET("/status/:storeid", statusHandler)
	router.Run(":8080")
}
