package main

import (
	"fmt"
    "os"
    "strconv"
)

func main(){
 if len(os.Args) != 2 {
        fmt.Println("Введите аргумент")
 		return
    }
    argument := os.Args[1]

	// с выражением после switch
    switch argument {
    case "0":
        fmt.Println("Ноль!")
    case "1":
        fmt.Println("Один!")
    case "2", "3", "4":
        fmt.Println("2 или 3 или 4")
        fallthrough
    default:
        fmt.Println("Значение:", argument)
 }

 value, err := strconv.Atoi(argument)
    if err != nil {
        fmt.Println("Не удается преобразовать в число:", argument)
 return
    }
	// без выражения после switch
    switch {
 case value == 0:
        fmt.Println("Ноль!")
    case value > 0:
        fmt.Println("Число позитивное")
    case value < 0:
        fmt.Println("Число негативное")
    default:
        fmt.Println("Непредвииденная ситуфция:", value)
 }
 }