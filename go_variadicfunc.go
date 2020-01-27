package main

import "fmt"

func find(num int, nums ...int)  {
	fmt.Printf("type of nums %T\n", nums)
	found := false
	for i, v := range nums{
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

func replace_word(s ...string)  {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main()  {
	nums := []int{42, 451, 52}
	find(89, 89, 90, 95)
	find(45, 67, 90, 45, 67)
	find(78, 85, 50, 123)
	find(78, nums...)

	welcome := []string{"hello", "lang"}
	replace_word(welcome...)
	fmt.Println(welcome)
}