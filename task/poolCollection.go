package task

func (c *PoolCollection) New() *PoolCollection {
	var poolCollection *PoolCollection = new(PoolCollection)
	poolCollection.Pool =  map[string]*Pool{}
	return poolCollection
}

func (c *PoolCollection) Add(nameChannels string, nameTask string, task *Task) {
	pool := Pool{Name: nameTask}
	pool.Add(nameTask,task)
	c.Pool[nameChannels] = &pool
}

