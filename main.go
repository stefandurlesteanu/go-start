package main

import "fmt"

func main() {
	// Array

	grades := [3]int{11, 12, 32}
	otherGrades := [...]int{1, 2, 3}
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Other grades: %v\n", otherGrades)
	students[0] = "Ion"
	students[1] = "Vasile"
	students[2] = "Adam"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students count: %v\n", len(students))

	a := [...]int{1, 2, 3}
	b := &a // Pointer to a value
	b[0] = 5
	fmt.Println(a)
	fmt.Println(b)

	//Slice

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   // slice of all elements
	e := c[3:]  // slice from 4th element to end
	f := c[:6]  // slice first 6 elements
	g := c[3:6] // slice the 4th, 5th, and 6th elements
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h := make([]int, 3, 100)

}
