package task

type HandlerFunc func(args map[string]string) error

type Task struct {
	Name  string            `json:"task"`
	Args  map[string]string `json:"args"`
	Retry int               `json:"retry"`
}