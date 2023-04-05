package main

import (
	service "assigment22/handler"
	control "assigment22/types"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleRequest(c *control.ItemSystem) {
	router := mux.NewRouter()
	handler := service.NewCollectionHandler(c)
	router.HandleFunc("/item/{itemName}", handler.SearchItemsByNameHandler).Methods("GET")       //added
	router.HandleFunc("/item/id/{itemId}", handler.GetItemByIdHandler).Methods("GET")            //added
	router.HandleFunc("/items/filter/price", handler.FilterItemsByPriceHandler).Methods("GET")   //added
	router.HandleFunc("/items/filter/rating", handler.FilterItemsByRatingHandler).Methods("GET") //added
	router.HandleFunc("/items", handler.GetItemsHandler).Methods("GET")                          //added
	router.HandleFunc("/users", handler.GetUserHandler).Methods("GET")                           //added
	router.HandleFunc("/item", handler.ItemPushHandler).Methods("POST")                          //added
	router.HandleFunc("/items/rating", handler.SetRatingHandler).Methods("POST")                 //added
	router.HandleFunc("/signin", handler.SignInHandler).Methods("POST")                          //added
	router.HandleFunc("/signup", handler.SignUpHandler).Methods("POST")                          //added
	router.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")                          //added

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	itemSystem := &control.ItemSystem{
		Users:  make(map[string]*control.User),
		Items:  make([]*control.Item, 0),
		ItemID: 0,
	}
	err := itemSystem.ReadFromFile()
	if err != nil {
		fmt.Println(err)
	}

	handleRequest(itemSystem)

}
