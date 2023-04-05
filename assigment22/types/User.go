package types

type User struct {
	Username string
	Password string
}

var IsAuthorized bool = false

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

type Item struct {
	ItmID  int
	Name   string
	Price  float64
	Rating float64
}

type ItemSystem struct {
	Users  map[string]*User
	Items  []*Item
	ItemID int
}

func NewItemSystem(users map[string]*User, items []*Item, itemID int) *ItemSystem {
	return &ItemSystem{Users: users, Items: items, ItemID: itemID}
}
