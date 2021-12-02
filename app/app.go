package app

import (
	"hexagonal/register"
)

type Notifyer interface {
	Notify(to, message string) error
}

type UserSaver interface {
	SaveUser(u register.User) error
}

type App struct {
	nf  Notifyer
	usv UserSaver
}

// ....
type User struct {
	Name     string `json:"name"`
	Password string `json:"psw"`
	Email    string `json:"email"`
}

func (a *App) RegisterUser(u User) error {
	// ....
	return register.RegisterUser(u.Name, u.Email, u.Password, a)
}

func (a *App) ProvideNewUser(u register.User) error {
	// ....
	if err := a.usv.SaveUser(u); err != nil {
		return err
	}
	return a.nf.Notify(u.Email, "Welcome to our service, "+u.Name)
}
