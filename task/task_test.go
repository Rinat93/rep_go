package task

import (
	"fmt"
	"testing"
)


func TestPool(t *testing.T) {
	t1 := Pool{Name:"test"}
	tasks := Task{
		Name:        "t1",
		Target:      "main",
		Time:        1111,
		CreatedTime: 1800,
		Save:        false,
		Type:        "",
		Prioritety:  0,
		QueueL:      nil,
		QueueR:      nil,
	}
	tasks2 := Task{
		Name:        "t12",
		Target:      "main",
		Time:        1111,
		CreatedTime: 1800,
		Save:        false,
		Type:        "",
		Prioritety:  2,
		QueueL:      nil,
		QueueR:      nil,
	}
	tasks3 := tasks2
	tasks3.Prioritety = 5

	t1.Add("task1",&tasks)
	t1.Add("task1",&tasks2)
	t1.Add("task1",&tasks2)
	t1.Add("task1",&tasks)
	t1.Add("task1",&tasks3)
	fmt.Printf("%+v\n" , t1.TaskGroup["task1"].QueueL)
}
