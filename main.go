package main

import (
	"log"
	"os"

	"ecommerce_go/database"
	//"ecommerce_go/middleware"
	"ecommerce_go/controllers"
	"ecommerce_go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.newApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-cart", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
