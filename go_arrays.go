package main

import "fmt"

func changeLocal(num [5]int)  {
	num[0] = 55
	fmt.Println("inside function", num)
}

func main() {
	num := [...]int{5, 4, 3, 2, 1}
	fmt.Println("before passing function", num)
	changeLocal(num)
	fmt.Println("after passing function", num)

	//a := [...]int{12, 54, 76}
	//fmt.Println(a)

	a := [...]float64{56.2, 12.5, 54.7, 56}
	fmt.Println("length of the array", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	sum := float64(0)
	for _, i := range a{
		sum += i
	}
	fmt.Println("\nsum of numbers in the list is", sum)

	s := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printArray(s)
	var b[3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "google"
	b[1][1] = "u-cell"
	b[2][0] = "beeline"
	b[2][1] = "mobiuz"
	fmt.Printf("\n")
	printArray(b)

	//c := [5]int{6, 7, 8, 2, 9}
	//var b []int = c[1:4]
	//fmt.Println(b)

	darr := [...]int{57, 52, 86, 23, 74, 96, 25}
	dsclice := darr[2:5]
	fmt.Println("array before", darr)

	for i := range dsclice{
		dsclice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{51, 42, 85}
	nums1 := numa[:]
	nums2 := numa[:]

	fmt.Println("array before change", numa)
	nums1[0] = 100
	fmt.Println("array after chaneg", nums1)
	nums2[1] = 101
	fmt.Println("array nums2 change", numa)

	fruitarr := [...]string{"apple", "orange", "peach", "melon"}
	fruitslice := fruitarr[1:3]
	fmt.Printf("length of the slice %d, capaicity %d", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("\nafter re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	i := make([]int, 5, 5)
	fmt.Println(i)

	cars := []string{"Ferrari", "Malibu", "Audi"}
	cars = append(cars, "Aston")
	fmt.Println(cars)

	var names []string
	if names == nil {
		fmt.Println("working...")
		names = append(names, "Azizbek", "Ahmad", "Mahmud")
		fmt.Println(names)
	}

	nos := []int{5, 3, 7}
	fmt.Println(nos)
	subtractOne(nos)
	fmt.Println(nos)

	pls := [][]string{
		{"C", "C++"},
		{"Python"},
		{"Go", "Java"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func printArray(s [3][2]string)  {
	for _,v1 := range s {
		for _, v2 := range v1 {
			fmt.Printf("%s ",v2)
		}
		fmt.Printf("\n")
	}
}

func subtractOne(numbers []int)  {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"US", "UZB", "UK", "ITA", "FRN"}
	neededContries := countries[:len(countries) - 2]
	contriesPy := make([]string, len(neededContries))
	copy(neededContries, contriesPy)
	return contriesPy
}