package types

import (
	"encoding/json"
	"os"
)

func (is *ItemSystem) Authenticate(username, password string) bool {
	for _, i := range is.Users {
		if i.Username == username && i.Password == password {
			return true
		}
	}
	return false
}
func (is *ItemSystem) UserCheck(username string) bool {
	for _, i := range is.Users {
		if i.Username == username {
			return true
		}
	}
	return false
}
func (is *ItemSystem) RegisterUser(username, password string) {
	user := NewUser(username, password)
	is.Users[username] = user
	err := is.UserSaveToFile()
	if err != nil {
		return
	}
}

func (is *ItemSystem) UserSaveToFile() error {
	file, err := os.Create("users.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(is.Users)
	if err != nil {
		return err
	}

	return nil
}
func (is *ItemSystem) GetUser(username string) []*User {
	var result []*User
	for _, user := range is.Users {
		if user.Username == username {
			result = append(result, user)
		}
	}
	return result
}
