package app

import "github.com/nhatnhanchiha/bookstore_items-api/controllers"

func mapUrl() {
	router.POST("/items", controllers.ItemsController.Create)
	router.GET("/ping", controllers.PingController.Ping)
	router.GET("/items/:id", controllers.ItemsController.Get)
	router.POST("/items/search", controllers.ItemsController.Search)
	router.PATCH("/items/:id", controllers.ItemsController.Patch)
	router.PUT("/items/:id", controllers.ItemsController.Put)
	router.DELETE("/items/:id", controllers.ItemsController.Delete)
}
