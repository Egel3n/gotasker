package task

import "fmt"

func Init() {
	Register("send_email", func(args map[string]string) error {
		to := args["to"]
		subject := args["subject"]
		fmt.Printf("[send_email] to: %s subject: %s\n", to, subject)
		return nil;
	})
}