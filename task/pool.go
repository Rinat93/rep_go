package task

import (
	"fmt"
)
func (c *Pool) New() *Pool {
	var pool *Pool = new(Pool)
	pool.TaskGroup = map[string]*Task{}
	return pool
}

/*
Add - Добавление задачи в Pool
	name - Имя группы задач
Добавляем задачу в Pool, в право кидаем задачи с боле высокой приоритетностью

 */
func (c *Pool) Add(name string, task *Task) error {
	if c.TaskGroup[name] == nil {
		c.TaskGroup[name] = task
	} else {
		if task == nil {
			return Error("Error: Not task")
		}
		c.TaskGroup[name].Add(task)
	}
	fmt.Printf("%+v\n", task)
	return nil
	//sort.Sort(c.TaskGroup)
}
