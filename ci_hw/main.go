// fix import cycle and make in works
package main

import (
	"ci_hw/course"
	"ci_hw/todo"
	"ci_hw/user"
)

func main() {
	u1 := user.NewUser(`Maks`, `Morozov`)
	td1 := todo.NewTODO(u1,
		`mon`, `Trim the cat`,
		`tue`, `Wash the car`,
		`wed`, `By cell phone`,
		`thu`, `Play guitar`,
		`fri`, `Go sleep`,
	)
	u1.SetTODO(td1)

	u2 := user.NewUser(`Alla`, `Pugachova`)
	td2 := todo.NewTODO(u2,
		`mon`, `Trim the cat`,
		`tue`, `Wash the car`,
		`wed`, `By cell phone`,
		`thu`, `Play guitar`,
		`fri`, `Go sleep`,
	)
	u2.SetTODO(td2)

	c := course.NewCourse(`Golang for beginners`, u1, u2)
	for _, u := range c.Users() {
		u.ShowMyTasks()
	}
}
