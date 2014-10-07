package controllers

import "github.com/revel/revel"

type User struct {
	Name     string
	Age      int
	Location string
}
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	user := User{
		Name:     "Pekka",
		Age:      16,
		Location: "Japan",
	}
	// vip := "abc"
	return c.Render(user)
}
