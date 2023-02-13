package main

import (
	c "assigment1/types"
	"fmt"
)

func main() {
	itemSystem := c.ItemSystem{
		Users:  make(map[string]*c.User),
		Items:  make([]*c.Item, 0),
		ItemID: 0,
	}
	err := itemSystem.ReadFromFile()
	if err != nil {
		fmt.Println(err)
	}
	var name, password string
	Authorized := false
	var operation int
	for true {
		fmt.Println("1)Add user\n2)Log In\n3)Log out\n4)Items Control\n5)Exit")
		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Print("Type a new username: ")
			fmt.Scanln(&name)
			fmt.Print("Type a new password: ")
			fmt.Scanln(&password)
			if itemSystem.Authenticate(name, password) || itemSystem.Reg(name) {
				fmt.Println("You're already registered")
			} else {
				itemSystem.RegisterUser(name, password)
			}

		} else if operation == 2 {
			fmt.Println("Type existing user to authorize")
			fmt.Print("Type a username: ")
			fmt.Scanln(&name)
			fmt.Print("Type a password: ")
			fmt.Scanln(&password)
			if itemSystem.Authenticate(name, password) {
				fmt.Println("Now, you can manage the items!")
				Authorized = true
			} else {
				fmt.Println("The data you have wrote is invalid! Try again.")
			}

		} else if operation == 3 {
			Authorized = false
			println("You're logged out")
		} else if operation == 4 {
			var OperOnItem int
			var itemName string
			var rating, price float64
			if Authorized {
				for true {
					fmt.Println("\n1)add item 2)filter item by price 3) filter item by rating, 4) search item by name 5) Give rating 6) exit items control")
					fmt.Print("Choose operation for item: ")
					fmt.Scanln(&OperOnItem)
					if OperOnItem == 1 {
						itemSystem.AddItem("item1", 100, 4)
						itemSystem.AddItem("item2", 200, 5)
						itemSystem.AddItem("item3", 50, 3)
						fmt.Print("Type a item name: ")
						fmt.Scanln(&itemName)
						fmt.Print("Type a rating: ")
						fmt.Scanln(&rating)
						fmt.Print("Type a price: ")
						fmt.Scanln(&price)
						itemSystem.AddItem(itemName, rating, price)
					} else if OperOnItem == 2 {
						fmt.Print("Type a price: ")
						fmt.Scanln(&price)
						fmt.Println(itemSystem.FilterItemsPrice(price))
					} else if OperOnItem == 3 {
						fmt.Print("Type a rating to filter: ")
						fmt.Scanln(&rating)
						fmt.Println(itemSystem.FilterItemsRating(rating))
					} else if OperOnItem == 4 {
						fmt.Print("Type a item name to search: ")
						fmt.Scanln(&itemName)
						filteredItems := itemSystem.SearchItems(itemName)
						for _, item := range filteredItems {
							fmt.Println(item.Name, item.Price, item.Rating)
						}
					} else if OperOnItem == 5 {
						fmt.Print("Type a item name: ")
						fmt.Scanln(&itemName)
						fmt.Print("Type a rating: ")
						fmt.Scanln(&rating)
						itemSystem.GiveRating(itemName, rating)
					} else if OperOnItem == 6 {
						break
					}
				}

			} else {
				println("You need to log in first")
			}
		} else if operation == 5 {
			err := itemSystem.UserSaveToFile()
			if err != nil {
				fmt.Println(err)
			}
			err = itemSystem.ItemSaveToFile()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("You're exit program")
			break
		} else {
			fmt.Println("Wrong Operation, try again.")
		}
	}
}
