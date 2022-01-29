package main

import (
	"log"
	"os"

	"github.com/anish700/ecommerce/database"

	// "github.com/anish700/ecommerce/middleware"

	"github.com/anish700/ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.newApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.new()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-cart", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
