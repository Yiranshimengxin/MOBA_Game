package log

import "fmt"

func Debug(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}
