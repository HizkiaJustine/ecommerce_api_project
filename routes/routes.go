package routes

import (
	"ecommerce_api_project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users/signup", controllers.SignUp())
	r.POST("/users/login", controllers.Login())
	r.GET("/users/productview", controllers.SearchProduct())
	r.GET("/users/search", controllers.SearchProductByQuery())
	r.POST("/admin/addproduct", controllers.ProductViewerAdmin())
}