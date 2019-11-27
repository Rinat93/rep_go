package task

/*
Task - Задачи
Name - название задачи
Target - Цель задачи(имя функции, метода и т.п.)
Time - время когда должно выполнится
CreatedTime - Время когда задача была создана
Save - Требуется ли сохранить задачу в файловую систему
Type - Тип задачи
	- 1 задача во все точки (om) после удаляется
	- 1 задача в 1 точку (oo) после удаляется
Prioritety - Приоритет задачи
*/

type Task struct {
	Name string
	Target string
	Time uint64
	CreatedTime string
	Save bool
	Type string
	Prioritety uint8
	QueueL *Task
	QueueR *Task
}

/*
Pool- Очередь задач
Name - Наименование потока, будет состоят из случайного хэша
Task - Список задач в пуле где map[string] это имя группы задач
*/
type Pool struct{
	Name string `json:"name"`
	TaskGroup map[string]*Task `json:"task_all"`
}

type PoolCollection struct {
	Pool []*Pool
}
