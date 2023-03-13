package logger

import "fmt"

func Log(a ...any) {
	fmt.Print("[tom] ")
	fmt.Println(a...)
}
