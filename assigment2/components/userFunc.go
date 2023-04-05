package components

//func (inv *UserCollection) Authenticate(username, password string) bool {
//	for _, i := range inv.users {
//		if i.Username == username && i.Password == password {
//			return true
//		}
//	}
//	return false
//}
//func (inv *UserCollection) UserCheck(username string) bool {
//	for _, i := range inv.users {
//		if i.Username == username {
//			return true
//		}
//	}
//	return false
//}
//func (inv *UserCollection) RegisterUser(username, password string) {
//	inv.userIterator++
//	user := NewUser(inv.userIterator, username, password)
//	inv.users = append(inv.users, user)
//
//}
//
//func (inv *UserCollection) UserSaveToFile() error {
//	file, err := os.Create("users.txt")
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	encoder := json.NewEncoder(file)
//	err = encoder.Encode(inv.users)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (inv *Inventory) CreateUser(username, password string) error {
	stmt, err := inv.db.Prepare("insert into users(username, password) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		return err
	}
	return nil
}
func (inv *Inventory) GetUserByUsername(username string) (*User, error) {
	var user User
	err := inv.db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).
		Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
