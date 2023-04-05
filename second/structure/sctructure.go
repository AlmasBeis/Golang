package structure

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser(username, password, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

type Product struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating int     `json:"rating"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
