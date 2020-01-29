package main

import "fmt"

//declaring struct
type Employee struct {
	fistName, lastName string
	age, salary int
}

//anonymous struct
type employee struct{
	first_name, last_name string
}

//nested struct
type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	address Address
}

func main()  {

	//creating struct using field names
	emp1 := Employee{
		fistName: "Azizbek",
		lastName: "Gaybullaev",
		age:      23,
		salary:   2500,
	}
	fmt.Println(emp1)

	//creating struct without field names
	emp2 := &Employee{"Mamura", "Rakhimjanova", 20, 2500}
	fmt.Println("First name: ", emp2.fistName)
	fmt.Println("Age: ", (*emp2).age)

	//creating anonymous struct
	emp3 := struct {
		first_name, last_name string
		age, salary int
	}{
		first_name: "Firdavs",
		last_name: "Olimov yoki Olimjanov",
		age: 0,
		salary: 0,
	}
	fmt.Println(emp3)

	//nested struct
	var p Person
	p.name = "Eshmat"
	p.age = 41
	p.address = Address{
		city: "San Fransisco",
		state: "California",
	}
	fmt.Println(p.address.city)

	//struct equality
	n1 := employee{"Steve", "Jobs"}
	n2 := employee{"Steve", "Jobs"}
	if n1 == n2 {
		fmt.Println("two struct are equal")
	}else {
		fmt.Println("are not equal")
	}
}