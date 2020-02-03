package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type myStr string

func (ms myStr) FindVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'o' || rune == 'u' || rune == 'i' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//practical use of interface
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId, basicPay, pf int
}

type Contract struct {
	empId, basicPay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator)  {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("\ntotal expense of month is $%d", expense)
}

//interface internal representation
type Tester interface {
	Test()
}

type myFloat float64

func (m myFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("\ninterface type %T value %v\n", t, t)
}

//empty interface
func define(i interface{})  {
	fmt.Printf("\nType=%T, value=%v", i, i)
}

//type assertion
func assert(i interface{})  {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

//type switch
func findType(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Printf("String and value is %s\n", i.(string))
	case int:
		fmt.Printf("integer and value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

type Describer interface {
	Describe()
}
type Person1 struct {
	name string
	age int
}

func (p Person1) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func find_type(i interface{})  {
	switch p := i.(type) {
	case Describer:
		p.Describe()
	default:
		fmt.Printf("unknoen person\n")
	}
}

func main()  {
	name := myStr("Azizbek")
	var v VowelsFinder
	v = name
	fmt.Printf("vowels are %c", v.FindVowels())

	//practical use of interface
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	//Interface internal representation
	var t Tester
	f := myFloat(97.5)
	t =f
	describe(t)
	t.Test()

	//empty interface
	s := "Hello world"
	define(s)
	i := 44
	define(i)
	strt := struct {
		name string
	}{
		name: "Azizbek Gaybullae",
	}
	define(strt)

	//type assertion
	var it interface{} = 56
	assert(it)
	var o_k interface{} = "bla bla bla"
	assert(o_k)

	//type switch
	findType("Azizbek")
	findType(41)
	findType(41.41)

	p := Person1{
		name: "Azizbek",
		age:  23,
	}
	find_type(p)
}
