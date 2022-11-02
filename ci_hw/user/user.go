package user

import (
	"ci_hw/course"
	"ci_hw/todo"
)

type User struct {
	name  string
	email string
	td    *todo.TODO
	on    *course.Course
}

func (u *User) ShowMyTasks() {
	u.td.ShowTasks()
}

func (u *User) SetTODO(td *todo.TODO) {
	u.td = td
}

func (u *User) SetCourse(c *course.Course) {
	u.on = c
}

func NewUser(n, e string) *User {
	return &User{name: n, email: e}
}
