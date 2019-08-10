package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	number1 := js.Global().Get("document").Call("getElementById", "number1").Get("value").String()
	number2 := js.Global().Get("document").Call("getElementById", "number2").Get("value").String()

	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	// result := js.ValueOf(i[0].Int() + i[1].Int())

	// result := number1 + number2
	result := int1 + int2
	fmt.Println("HASIL")
	fmt.Printf("RESULT %#v %#v %#v", number1, number2, result)

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
