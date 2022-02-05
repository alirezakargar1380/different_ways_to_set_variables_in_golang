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

	// second way to declaration a variable -> short decalaration
	what_are_you_doing := "learning Golang"
	fmt.Println("hello world", what_are_you_doing)

	// short declaration multiple variables
	car, cost := "Bmw", 5999
	fmt.Println("car:", car, "cost:", cost)

	// update a declared variable
	// car, cost := "Bmw", 5999 --> it give you an error

	var oppend bool = false
	oppend, time := true, 18
	fmt.Println("status:", oppend, "time:", time)

	// normal declaration multiple variables
	var (
		job    string
		salary float32
	)
	job, salary = "programmer", 156.025

	fmt.Println("what is your job?", job)
	fmt.Println("how much you get salary?", salary)

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
}
