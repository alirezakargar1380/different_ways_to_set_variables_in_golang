package main

import (
	"fmt"
)

func main() {
	// keyword variable_name type
	var car_name string = "ford"
	var car_model string = "mustang"
	var car_number int = 1992
	// tip: when you declaration a variable, you should use it.
	// go gives you an error for unsued variable

	_ = car_number
	// if you don't want to get compile error -> fist way
	_, _ = car_name, car_model
	// if you don't want to get compile error -> second way

	// second way to declaration a variable
	what_are_you_doing := "learning Golang"
	fmt.Println("hello world", what_are_you_doing)

	// declaration multiple variables
	car, cost := "Bmw", 5999
	fmt.Println("car:", car, "cost:", cost)
}
