package task

//import (
//	"sort"
//)
func (c *Pool) New() *Pool {
	var pool *Pool = new(Pool)
	pool.TaskGroup = map[string]*Task{}
	return pool
}

func (c *Pool) Add(name string, task *Task) {
	if c.TaskGroup[name] == nil {
		c.TaskGroup = map[string]*Task{}
		c.TaskGroup[name] = task
	} else {
		if task.Prioritety > c.TaskGroup[name].Prioritety{
			c.TaskGroup[name].QueueR = task
			c.Add(name, task.QueueL)
		} else {
			c.TaskGroup[name].QueueL = task
			c.Add(name, task.QueueR)
		}

	}
	//sort.Sort(c.TaskGroup)
}

