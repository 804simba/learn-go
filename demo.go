package main

import "fmt"

func main() {
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
}

func test(array [3][2]int) {
	array[0] = [2]int{100, 200}
}
