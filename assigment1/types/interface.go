package types

type UserInterface interface {
	Authenticate(username, password string) bool
	UserCheck(username string) bool
	RegisterUser(username, password string)
	UserSaveToFile() error
}
type ItemInterface interface {
	SearchItems(name string) []*Item
	FilterItemsPrice(price float64) []*Item
	FilterItemsRating(rating float64) []*Item
	AddItem(name string, price, rating float64)
	GiveRating(itemName string, rating float64) bool
	ItemSaveToFile() error
}
