package main

import "fmt"

//pointers implementation
type Describer1 interface {
	Describe()
}

type Person_1 struct {
	name string
	age  int
}

func (p Person_1) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address1 struct {
	city, state string
}

func (a *Address1) Describe() {
	fmt.Printf("%s state %s country", a.city, a.state)
}

//implementing multiple interfaces
type SalaryCal interface {
	ShowSalary()
}

type LeaveCalculator interface {
	CalculatorLeavesLeft() int
}

//embedding interfaces
type EmployeeOperations interface {
	SalaryCal
	LeaveCalculator
}

type Employeee struct {
	firstName, lastName      string
	basicPay, pf             int
	totalLeaves, leavesTaken int
}

func (e Employeee) ShowSalary() {
	fmt.Printf("\n%s %s has salary $%d", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employeee) CalculatorLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	var d1 Describer1
	p1 := Person_1{"Azizbek", 23}
	d1 = p1
	d1.Describe()
	p2 := Person_1{"james", 24}
	d1 = &p2
	d1.Describe()

	var d2 Describer1
	a := Address1{
		city:  "San Fransisco",
		state: "California",
	}
	d2 = &a
	d2.Describe()

	e := Employeee{
		firstName:   "Azizbek",
		lastName:    "Gaybullaev",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	//var s SalaryCal = e
	//s.ShowSalary()
	//var l LeaveCalculator = e
	//fmt.Println("\nLeaves left =", l.CalculatorLeavesLeft())

	var empOp EmployeeOperations = e
	empOp.ShowSalary()
	fmt.Println("\nLeaves left =", empOp.CalculatorLeavesLeft())
}
