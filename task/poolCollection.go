package task

func (c *PoolCollection) New() *PoolCollection {
	var poolCollection *PoolCollection = new(PoolCollection)
	return poolCollection
}

/*
Add добавление
namePool имя пула который следует добавить в коллекцию
nameTaskGroup - Имя группы тасков
task - Массив задач который должен будет выполнить данная коллекция пулов
*/
func (c *PoolCollection) Add(namePool string, nameTaskGroup string, task []*Task) {
	pool := new(Pool).New()
	pool.Name = namePool
	for _,t := range task {
		pool.Add(nameTaskGroup, t)
	}
	c.Pool = append(c.Pool, pool)
}

