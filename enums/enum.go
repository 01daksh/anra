package enums

const (
	Todo       string = "todo"
	InProgress string = "in_progress"
	Done       string = "done"
)

var statusMap = map[string]bool{
	Todo:       true,
	InProgress: true,
	Done:       true,
}

func IsStatusValid(status string) bool {
	return statusMap[status]
}
