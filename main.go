package main

import (
	"ecommerce_api_project/controllers"
	"ecommerce_api_project/database"
	"ecommerce_api_project/middleware"
	"ecommerce_api_project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(
		database.ProductData(database.Client, "Products"), 
		database.UserData(database.Client, "Users"),
	)

	r := gin.New()
	r.Use(gin.Logger())

	routes.UserRoutes(r)
	r.Use(middlewares.Authentication())

	r.GET("/addtocart", app.AddToCart())
	r.GET("/removeitem", app.RemoveItem())
	r.GET("/cartcheckout", app.Checkout())
	r.GET("/instantbuy", app.InstantBuy())

	log.Fatal(r.Run(":" + port))
}