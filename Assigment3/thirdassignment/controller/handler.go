package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"thirdassignment/model"
)

type bodyHeader struct {
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

func (c *Controller) getAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	books, err := c.service.BookLibrary.GetAllBook()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := c.service.BookLibrary.GetBookById(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) getBooksByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var body bodyHeader
	err := json.NewDecoder(r.Body).Decode(&body)
	//fmt.Println(body.Title)
	books, err := c.service.BookLibrary.GetBooksByTitle(body.Title)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)

}

func (c *Controller) getBooksOrderedByCost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var body bodyHeader
	err := json.NewDecoder(r.Body).Decode(&body)
	//fmt.Println(body.Sort)
	books, err := c.service.BookLibrary.GetOrderedBooksByCost(body.Sort)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)

}

func (c *Controller) addBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.service.BookLibrary.AddBook(book); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) updateBookById(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.service.BookLibrary.SetBook(book, id); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (c *Controller) deleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = c.service.BookLibrary.RemoveBook(id); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
