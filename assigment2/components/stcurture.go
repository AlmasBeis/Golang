package components

import "database/sql"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(id int, username string, password string) *User {
	return &User{Id: id, Username: username, Password: password}
}

type Item struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}

func NewItem(id int, name string, description string, price float64, rating float64) *Item {
	return &Item{Id: id, Name: name, Description: description, Price: price, Rating: rating}
}

type Collection struct {
	userIterator, itemIterator int
	Users                      []User
	Items                      []Item
}

func NewCollection(userIterator int, itemIterator int, users []User, items []Item) *Collection {
	return &Collection{userIterator: userIterator, itemIterator: itemIterator, Users: users, Items: items}
}

type Inventory struct {
	db *sql.DB
}

func NewInventory(db *sql.DB) *Inventory {
	return &Inventory{db: db}
}
