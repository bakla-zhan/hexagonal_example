package register

type User struct {
	Name     string
	Password string
	Email    string
}

type UserProvider interface {
	ProvideNewUser(u User) error
}

func RegisterUser(name, email, pass string, up UserProvider) error {
	u := User{name, pass, email}
	// ....
	return up.ProvideNewUser(u)
}
