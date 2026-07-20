package tasks

import "fmt"

func Info() {
	fmt.Println("")
}

func PrintTask(name, link string) {
	fmt.Println(name)
	fmt.Println(link)
}

func PrintScript[T any, R any](fn func(T) R, arg T) {
	fmt.Println("")
	fmt.Println("Input:", arg)
	fmt.Println("Output:", fn(arg))
}

func PrintScriptTwo[A, B any, R any](fn func(A, B) R, a A, b B) {
	fmt.Println("")
	fmt.Println("Input:", a, b)
	fmt.Println("Output:", fn(a, b))
}
