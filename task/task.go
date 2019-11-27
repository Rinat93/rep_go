package task

import (
	"time"
)
func (t *Task) New() *Task {
	var file *Task = &Task{
		Name:        "",
		Target:      "",
		Time:        0,
		CreatedTime: time.RFC822,
		Save:        false,
		Type:        "",
		Prioritety:  0,
		QueueL:      nil,
		QueueR:      nil,
	}
	return file
}

func (t *Task) Add(node *Task) {
	if node == nil {
		return
	}

	if t.Name == "" {
		t = node
		return
	}
	if t.Prioritety < node.Prioritety {
		if t.QueueR == nil {
			t.QueueR = new(Task).New()
			t.QueueR.Add(node)
		} else {
			if t.QueueL == nil {
				t.QueueL = new(Task).New()
			}
			t.QueueL.Add(node)
		}
	}
}
