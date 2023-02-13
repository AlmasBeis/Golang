package types

func (u *User) Login(username, password string) bool {
	if username == u.Username && password == u.Password {
		return true
	}
	return false
}
func (u *User) RegistrationCheck(username string) bool {
	if username == u.Username {
		return true
	}
	return false
}
