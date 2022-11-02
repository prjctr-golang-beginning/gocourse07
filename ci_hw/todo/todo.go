package todo

import (
	"ci_hw/user"
	"fmt"
)

type TODO struct {
	owner *user.User
	tasks []Task
}

func (t TODO) ShowTasks() {
	for _, task := range t.tasks {
		fmt.Println(task)
	}
}

func (t *TODO) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

type Task struct {
	name        string
	description string
}

func (t Task) String() string {
	return fmt.Sprintf(`name: %s, description: %s`, t.name, t.description)
}

// NewTODO expects name, description, name, description and else...
func NewTODO(o *user.User, text ...string) *TODO {
	if len(text)%2 != 0 {
		panic(`arg must be even`)
	}

	res := &TODO{owner: o}

	var i int
	for {
		res.AddTask(Task{text[i], text[i+1]})
		i += 2
		if i >= len(text) {
			break
		}
	}

	return res
}
