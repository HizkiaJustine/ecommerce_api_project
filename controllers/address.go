package controllers

import (
	"context"
	"ecommerce_api_project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")

		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
			c.Abort()
			return 
		}

		addresses := make([]models.Address, 0)
		usertID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertID}}
		update := bson.D{{ Key: "$set", Value: bson.D{primitive.E{Key: "addresses", Value: addresses}} }}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "address not found")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "address deleted successfully")
	}
}
