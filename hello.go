package hello

import(
	"fmt"
)

func Hello(msg string) string{
	fmt.Printf("msg: %s", msg)
	return fmt.Sprint("Hello, %s", msg)
}
