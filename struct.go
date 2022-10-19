package main

import "fmt"

type Student struct {
	name  string
	grade int
}

func newStudent() *Student {
	return &Student{
		name:  "abdil",
		grade: 1,
	}
}

func main() {
	var s1 Student
	s1 = *newStudent()
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)
}
