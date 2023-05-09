package routes

import (
	"github.com/dhwanidoshi07/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/getAll", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}/getBookId", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}/update", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}/delete", controllers.DeleteBook).Methods("DELETE")
}
