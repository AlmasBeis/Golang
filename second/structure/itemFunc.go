package structure

import "database/sql"

func SearchProducts(name string) ([]*Product, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Search for products by name
	rows, err := db.Query("SELECT id, name, price, rating FROM products WHERE name LIKE '%' || $1 || '%'", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Convert the
	var products []*Product
	for rows.Next() {
		p := &Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Rating)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
func FilterProducts(minPrice, maxPrice float64, minRating int) ([]*Product, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Filter products by price and rating
	rows, err := db.Query("SELECT id, name, price, rating FROM products WHERE price >= $1 AND price <= $2 AND rating >= $3", minPrice, maxPrice, minRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Convert the results into Product instances
	var products []*Product
	for rows.Next() {
		p := &Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Rating)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (p *Product) RateProduct(rating int) error {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb")
	if err != nil {
		return err
	}
	defer db.Close()

	// Update the product's rating in the database
	_, err = db.Exec("UPDATE products SET rating = $1 WHERE id = $2", rating, p.ID)
	if err != nil {
		return err
	}

	// Update the product's rating in the Product instance
	p.Rating = rating

	return nil
}
