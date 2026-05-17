package main

import "fmt"

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	// defer works as finally block in Java
	// anything after a panic statement would not run
	// recover is used to catch the panic and continue execution (it is always put in a deferred function)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result, err := divide(10, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", result)
	fmt.Println("Hello, World!")

	resultOfAddition := addNumbers(10, 20)
	fmt.Println("Result of addition:", resultOfAddition)

	resultOfFloatAddition := addNumbers(10.5, 20.3)
	fmt.Println("Result of float addition:", resultOfFloatAddition)
}

// Generics

type Number interface {
	int | float64 | uint
}

func addNumbers[T Number](x T, y T) T {
	return x + y
}

type GenericSlice[T any] []T

func (s GenericSlice[T]) Len() int {
	return len(s)
}

type GenericStruct[T any, K any] struct {
	key T
	val K
}
