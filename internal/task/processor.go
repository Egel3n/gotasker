package task

import "errors"

func Process(task *Task) error {

	handler, ok := GetHandler(task.Name)

	if !ok {
		return errors.New("No handler registered for this task name: "+ task.Name)
	}

	err := handler(task.Args)
	if err != nil {
		return err
	}

	return nil
}