package controllers

import (
	"context"
	"ecommerce_api_project/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid code"})
			c.Abort()
			return

		}
		address, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var addresses models.Address

		addresses.AddressID = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		matchFilter := bson.D{{ Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}} }}
		unwind := bson.D{{ Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}} }}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}} }}
		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var addressInfo []bson.M
		if err = pointCursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32
		for _, addressNo := range addressInfo {
			count := addressNo["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{ Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}} }}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			c.IndentedJSON(400, "Not Allowed")
		}
		defer cancel()
		ctx.Done()
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}
		usertID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var editAddress models.Address
		if err = c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertID}}
		update := bson.D{{ Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house", Value: editAddress.House}, {Key: "address.0.street", Value: editAddress.Street}, {Key: "address.0.city", Value: editAddress.City}}} }
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "address updated successfully")
	}

}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}
		usertID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var editAddress models.Address
		if err = c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertID}}
		update := bson.D{{ Key: "$set", Value: bson.D{primitive.E{Key: "address.1.house", Value: editAddress.House}, {Key: "address.1.street", Value: editAddress.Street}, {Key: "address.1.city", Value: editAddress.City}}} }
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "address updated successfully")
	}

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
