package hello

import(
	"fmt"
)

func Hello(msg string) string{
	return fmt.Sprint("Hello, %s", msg)
}
