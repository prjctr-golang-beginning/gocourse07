package course

import (
	"ci_hw/user"
)

type Course struct {
	title string
	users []*user.User
}

func (c *Course) addUser(u *user.User) {
	c.users = append(c.users, u)
}

func (c *Course) Users() []*user.User {
	return c.users
}

func NewCourse(title string, u ...*user.User) *Course {
	c := &Course{title: title}
	for i := range u {
		c.addUser(u[i])
	}

	return c
}
