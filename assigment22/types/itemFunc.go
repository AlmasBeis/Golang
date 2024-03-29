package types

import (
	"encoding/json"
	"os"
)

func (is *ItemSystem) SearchItems(name string) []*Item {
	var results []*Item
	for _, item := range is.Items {
		if name == item.Name {
			results = append(results, item)
		}
	}
	return results
}

func (is *ItemSystem) FilterItemsPrice(price float64) []*Item {
	var filtered []*Item
	for _, item := range is.Items {
		if price == 0 || price >= item.Price {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
func (is *ItemSystem) FilterItemsRating(rating float64) []*Item {
	var filtered []*Item
	for _, item := range is.Items {
		if rating == 0 || rating <= item.Rating {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func (is *ItemSystem) AddItem(name string, price, rating float64) {
	is.ItemID++
	item := &Item{is.ItemID, name, price, rating}
	is.Items = append(is.Items, item)
	err := is.ItemSaveToFile()
	if err != nil {
		return
	}
}

func (is *ItemSystem) GiveRating(itemName string, rating float64) bool {
	for _, i := range is.Items {
		if itemName == i.Name {
			i.Rating = rating
		}
	}
	return true
}

func (is *ItemSystem) ItemSaveToFile() error {
	file, err := os.Create("items.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var encoder = json.NewEncoder(file)
	err = encoder.Encode(is.Items)
	if err != nil {
		return err
	}
	return nil
}
func (is *ItemSystem) ReadFromFile() error {
	file, err := os.Open("items.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&is.Items)
	if err != nil {
		return err
	}

	file, err = os.Open("users.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder = json.NewDecoder(file)
	err = decoder.Decode(&is.Users)
	if err != nil {
		return err
	}

	return nil
}
