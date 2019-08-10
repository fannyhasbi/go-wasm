package main

import (
	"fmt"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() + i[1].Int())
	fmt.Println("RESULT", result)

	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM initialized")

	registerCallbacks()
	<-c
}
