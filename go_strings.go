package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string)  {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
}

func printChars(s string)  {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string)  {
	for index, runes := range s{
		fmt.Printf("%c starts at byte %d\n", runes, index)
	}
}

func length(s string)  {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main()  {
	name := "Hello world"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	printCharsAndBytes(name)

	byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)
	fmt.Println(len(str))
	length(str)

	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
