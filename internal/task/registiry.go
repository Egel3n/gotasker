package task

var handlers = make(map[string]HandlerFunc)

func Register(name string, fn HandlerFunc) {
	handlers[name] = fn
}

func GetHandler(name string) (HandlerFunc, bool) {
	fn, ok := handlers[name]
	return fn, ok
}