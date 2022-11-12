package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.CreateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.CreateBook).Methods("DELETE")
}
