package task

import "testing"

// Генерируем тестовые данные
func GeneratePool(namePool string, nameGroupTask string) (string, string, []*Task) {
	var tasks []*Task

	for i := 0; i < 15; i++ {
		t := new(Task).New()
		t.Name = "test"+string(i)
		t.Prioritety = uint8(i)
		tasks = append(tasks, t)
	}
	return namePool, nameGroupTask,tasks
}


func TestPoolCollection(t *testing.T) {
	poolCol := new(PoolCollection)
	poolCol.Add(GeneratePool("Test1","TaskGroup1"))
	poolCol.Add(GeneratePool("Test1","TaskGroup2"))
	poolCol.Add(GeneratePool("Test1","TaskGroup3"))
	poolCol.Add(GeneratePool("Test1","TaskGroup4"))
}

