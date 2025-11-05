package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productscontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//Products
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/products/add", productscontroller.Add)
	http.HandleFunc("/products/edit", productscontroller.Edit)
	http.HandleFunc("/products/delete", productscontroller.Delete)
	http.HandleFunc("/products/detail", productscontroller.Detail)

	log.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)
}
