package components

func (inv *Inventory) AddItem(item *Item) error {
	stmt, err := inv.db.Prepare("INSERT INTO items(name, description, price, rating) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Name, item.Description, item.Price, item.Rating)
	if err != nil {
		return err
	}

	return nil
}

func (inv *Inventory) GetItemById(itemId int) (*Item, error) {
	var item Item
	err := inv.db.QueryRow("SELECT id, name, description, price, rating FROM items WHERE id=?", itemId).
		Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.Rating)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (inv *Inventory) SearchItemsByName(name string) ([]*Item, error) {
	likeStr := "%" + name + "%"
	rows, err := inv.db.Query("SELECT id, name, description, price, rating FROM items WHERE name LIKE ?", likeStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.Rating)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (inv *Inventory) FilterItemsByPriceRange(minPrice, maxPrice float64) ([]*Item, error) {
	rows, err := inv.db.Query("SELECT id, name, description, price, rating FROM items WHERE price BETWEEN ? AND ?", minPrice, maxPrice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.Rating)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (inv *Inventory) FilterItemsByRatingRange(minRating, maxRating float64) ([]*Item, error) {
	rows, err := inv.db.Query("SELECT id, name, description, price, rating FROM items WHERE rating BETWEEN ? AND ?", minRating, maxRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.Rating)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
func (inv *Inventory) RateItem(itemId int, rating float64) error {
	stmt, err := inv.db.Prepare("UPDATE items SET rating=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(rating, itemId)
	if err != nil {
		return err
	}

	return nil
}
