package models

// going to use the in-memory storage, otherwise would have gone with gorm tags
type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskCollection struct {
	Tasks []Task
}

func NewTaskCollection() *TaskCollection {
	return &TaskCollection{
		Tasks: []Task{},
	}
}

func (collection *TaskCollection) AddTask(task Task) {
	collection.Tasks = append(collection.Tasks, task)
}

func (collection *TaskCollection) GetTasks() []Task {
	return collection.Tasks
}
