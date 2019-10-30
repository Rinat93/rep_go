package task

//import (
//	"sort"
//)
func (c *Pool) New() *Pool {
	var pool *Pool = new(Pool)
	pool.TaskGroup = map[string]*Task{}
	return pool
}

/*
Add - Добавление задачи в Pool

Добавляем задачу в Pool, в право кидаем задачи с боле высокой приоритетностью

 */
func (c *Pool) Add(name string, task *Task) {
	if c.TaskGroup[name] == nil {
		c.TaskGroup = map[string]*Task{}
		c.TaskGroup[name] = task
	} else {
		if task == nil {
			return
		}
		c.TaskGroup[name].Add(task)

	}
	//sort.Sort(c.TaskGroup)
}
