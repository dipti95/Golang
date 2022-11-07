package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	make string
	year uint16
	used bool
}

type Student struct{
	id int
	name string
}

type Classroom struct{
	id int
	capacity int
	subject string
	studentList []Student
}

func (c Car) describe() string {

	used := ""
	if c.used {
		used ="a used car"
	}else {
		used ="a new car"
	}
  	return "This " + strconv.Itoa(int(c.year)) +" "+c.make+" is " + used+"."

} 

func main (){


	car1 := Car{make:"Delorean", year:1985, used: true}
	car2 := Car{make:"Honda", year:2023, used :false}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.describe())

	fmt.Println(car2.describe())

  // Exercise

  c1 := Classroom{id:1, capacity:10, subject:"Maths", studentList: []Student {
	{
		id :2,
		name:"Tim",
	},
	{
		id:9,
		name:"Tom",
	},
  }}

  c2 := new (Classroom)

  c2.id = 2
  c2.capacity = 40
  c2.subject = "Chemistry"
  c2.studentList = []Student{
	{
		id:10,
		name:"Jeni",
	},
	{
		id:30,
		name:"Max",
	},
  }

  fmt.Println(c1)
  fmt.Println(c2)
}