package controller

import (
	"github.com/gorilla/mux"
	"thirdassignment/service"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", c.getAllBook).Methods("GET")
	router.HandleFunc("/books/title", c.getBooksByTitle).Methods("GET")
	router.HandleFunc("/books/cost", c.getBooksOrderedByCost).Methods("GET")
	router.HandleFunc("/books/{id}", c.getBookById).Methods("GET")
	router.HandleFunc("/books", c.addBook).Methods("POST")
	router.HandleFunc("/books/{id}", c.updateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", c.deleteBookById).Methods("DELETE")

	return router
}
