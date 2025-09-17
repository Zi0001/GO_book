package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("Первый символ", string(aString[0]))

	// руны
    // руна
	r := '€'
	fmt.Println("значение int32:", r)
	// преобразование рун в текст
	fmt.Printf("строка: %s символ: %c\n", r, r)

	// вывести существующую строку в виде рун
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// String to rune Array
	// myRune := []rune(aString)
	// fmt.Printf("myRune %U\n", myRune)

	// Rune array to string
	// runeArray := []rune{'1', '2', '3'}
	// s := string(runeArray)
	// fmt.Println(s)

	// вывести существующую строку в виде символов
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}