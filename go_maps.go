package main

import "fmt"

/*
	A map is a builtin type in Go which associates a value to a key.
	The value can be retrieved using the corresponding key.
*/

func main()  {
	var employeeSalary = map[string]int{
		"Azizbek": 145200,
		"Doston": 41000,
	}
	//if employeeSalary == nil {
	//	fmt.Println("map is nil. Going to make one")
	//	employeeSalary = make(map[string]int)
	//}

	employeeSalary["Mike"] = 4120
	employee := "joe"
	value, ok := employeeSalary[employee]

	if ok == true {
		fmt.Println("salary of employee", employee, "is", value)
	} else {
		fmt.Println(employee, "not found")
	}
	fmt.Println("original person salary", employeeSalary)
	//newPersonSalary := employeeSalary
	//newPersonSalary["Mike"] = 12540
	//fmt.Println("salary changed", employeeSalary)
	//delete(employeeSalary, "Mike")
	//for key, values := range employeeSalary {
	//	fmt.Printf("employeeSalary[%s] = %d\n", key, values)
	//}

	//fmt.Println("length of map is", len(employeeSalary))

	CEOSalary := map[string]int{
		"Azizbek": 145200,
		"Doston": 41000,
		"Mike": 4121,
	}
		for _, v := range employeeSalary {
			for _, v2 := range CEOSalary {
				if v == v2 {
					fmt.Println(v)
					fmt.Println(true)
				}else {
					fmt.Println(false)
				}
			}
		}
}