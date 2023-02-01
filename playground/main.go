package main

import (
	"fmt"
	"playground/custommath"
	"playground/pyramid"
	"playground/user"
)

func main() {
	var firstName string = "sorrawit"
	var lastName string = "kwanja"
	var age uint = 20
	var message string = fmt.Sprintf("My name is %v %v. I am %v years old.", firstName, lastName, age)
	fmt.Println(message)

	fmt.Println(custommath.Add(5, 10))

	john := user.Student{
		Name:  "John Doe",
		Year:  3,
		Major: "computer science",
	}
	fmt.Println(john)
	fmt.Println(john.Major)
	fmt.Println(john.Year)
	fmt.Println(john.Name)
	fmt.Printf("type of %v is %T\n", john.Name, john.Name)
	// fmt.Println(john.Greeting()) <- error (no values) used as value
	// happend when calling some function and pass in another function
	// but the second function does not return anything (void)
	john.Greeting()

	pyramid.Pyramid()

	arr := []user.Student{john}
	fmt.Println(arr)

	country := map[string]string{}
	country["TH"] = "Thailand"
	country["JP"] = "Japan"
	country["EN"] = "English"
	fmt.Println(country["TH"])
}

// basic function syntax
// func addTwoNumber(x int, y int) int {
//     return x + y
// }
// function can return multiple values
// func divideTwoNumber(x int, y int) (int, error) {
//     if y == 0 {
//         return 0, errors.New("Cannot divide by 0")
//     }
//     return x / y, nil
// }

// Loop
// arr := []int{1,2,3,4,5}
// loop with range (foreach)
// for i, v := range arr {
//     if i == len(arr) - 1 {
//         fmt.Println(v)
//     } else {
//         fmt.Print(v, " ")
//     }
// }
// loop with initializer
// for i := 0; i < len(arr); i++ {
//     if i == len(arr) - 1 {
//         fmt.Println(arr[i])
//     } else {
//         fmt.Print(arr[i], " ")
//     }
// }
// while loop
// var i uint = 0
// for i < 10 {
//     fmt.Println(i)
// }
// while(true) loop
// for {
//     fmt.Println("Infinite Loop")
// }

// Array and Slice
// Array:
// var arr1 [5]uint
// arr1[0] = 5
// arr1[1] = 1
// arr1[2] = 3

// arr2 := [5]uint
// arr2[0] = 3
// arr2[1] = 1
// arr2[2] = 5

// Slice (Dynamic array):
// append return a new array (immutable)
// var arr3 []uint
// arr3 = arr3.append(33)
// arr3 = arr3.append(99)
// arr3 = arr3.append(55)

// arr4 := []uint
// arr4 = arr4.append(88)
// arr4 = arr4.append(22)
// arr4 = arr4.append(11)

// syntax: slice[low:high]
// arr5 := arr4[:1] // get elements from 0 to 1
// arr6 := arr4[1:] // get elements from 1 to last
// arr7 := arr4[1:2] // get elements from 1 to 1
