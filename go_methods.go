package main

import (
	"fmt"
	"math"
)

type Employees struct {
	name, currency string
	salary int
}

type Employer struct {
	name string
	age int
}

//sample method
func (e Employees) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length, width int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//pointers with receivers
//method with value receiver
func (e Employer) changeName(newName string) {
	e.name = newName
}

//method with pointer receiver
func (e *Employer) changeAge(newAge int) {
	e.age = newAge
}

func main()  {
	emp1 := Employees{
		name:     "Azizbek",
		currency: "$",
		salary:   250000000,
	}
	fmt.Println(emp1)
	emp1.displaySalary()
	
	r := Rectangle{
		length: 15,
		width:  5,
	}
	fmt.Printf("\narea of rectangle %d\n", r.Area())

	c := Circle{radius:12}
	fmt.Printf("area of circle %f", c.Area())

	e := Employer{
		name: "Azizbek",
		age:  24,
	}
	fmt.Printf("employee name before change: %s", e.name)
	e.changeName("Azizbek")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployer age before age change %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nemployer age after change: %d", e.age)
}