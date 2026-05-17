package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) CalculateDateOfBirth() string {
	return time.Now().Format("2006-01-02")
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var areaOfTriangle Shape = Rectangle{Width: 10, Height: 5}
	fmt.Println(areaOfTriangle.Area())
	//fmt.Println("Hello World")
	//var name = "Timothy"
	//fmt.Println(name)
	//var age uint8 = 255
	//fmt.Println(age)
	//myClass := "level 200" // implicit defined type
	//fmt.Println(myClass)

	str := "-"

	for idx := 0; idx < len(str); idx++ {
		fmt.Printf("%c \n \n", str[idx])
	}

	fullName := "Timothy Olisaeloka Ngonadi"
	// loop through a string
	for _, char := range fullName {
		fmt.Printf("%c", char)
	}

	var array [2]bool
	array[0] = true
	array[1] = true
	fmt.Println(array)
	fmt.Println(len(array))

	myArray := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	// This passes a copy of the original array to the function,
	// it does not mutate the underlying array
	test(myArray)
	fmt.Println(myArray)

	var myslice1 []int
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	fmt.Println(myslice1, myslice2)

	sliceOfNames := []string{"Timothy", "Olisaeloka", "Timothy"}

	for x := 0; x < 10; x++ {
		sliceOfNames = append(sliceOfNames, "Timothy")
		fmt.Println(sliceOfNames, len(sliceOfNames), cap(sliceOfNames))
	}

	fmt.Println(sliceOfNames)

	dictionary := map[string]string{"name": "Timothy", "age": "200"}
	dictionary["class"] = "level 200"
	delete(dictionary, "age")
	fmt.Println(dictionary)
	value, ok := dictionary["class"]
	fmt.Println(value, ok)

	fmt.Println(add(10, 20))

	result := concatenate("trump")
	concat1 := result("et")
	concat2 := result("al")
	fmt.Println(concat1, concat2)

	person := Person{"Timothy", 20}
	fmt.Println(person.Name)
	fmt.Println(person.CalculateDateOfBirth())
}

func test(array [3][2]int) {
	array[0] = [2]int{100, 200}
}

func add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

func concatenate(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}
