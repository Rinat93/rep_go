package task

//import (
//	"sort"
//)
func (c *Pool) New() *Pool {
	var pool *Pool = new(Pool)
	pool.TaskGroup = map[string][]*Task{}
	return pool
}

func (c *Pool) Add(name string, task *Task) {
	c.TaskGroup[name] = append(c.TaskGroup[name], task)
	//sort.Sort(c.TaskGroup)
}

