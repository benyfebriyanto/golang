package main

import (
	"fmt"
	"log"
	"mod2-http-method/controllers"
	"mod2-http-method/services"
	"net/http"
)

func main() {
	// Initialize the service and controller
	itemService := services.ItemService{}
	itemController := controllers.ItemController{ItemService: itemService}

	// Define routes and attach the handler functions
	http.HandleFunc("/items", itemController.GetAllItemsHandler)
	http.HandleFunc("/item", itemController.GetItemHandler)
	http.HandleFunc("/item/create", itemController.CreateItemHandler)
	http.HandleFunc("/item/update", itemController.UpdateItemHandler)
	http.HandleFunc("/item/delete", itemController.DeleteItemHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
