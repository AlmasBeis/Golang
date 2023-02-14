package types

type User struct {
	Username string
	Password string
}

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

type Item struct {
	itmID  int
	Name   string
	Price  float64
	Rating float64
}

type ItemSystem struct {
	Users  map[string]*User
	Items  []*Item
	ItemID int
}
