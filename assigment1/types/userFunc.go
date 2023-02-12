package types

func (u *User) Login(username, password string) bool {
	if username == u.Username && password == u.Password {
		return true
	}
	return false
}
