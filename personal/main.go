package main

import (
	c "personal/converter"
	"fmt"
)


func main(){
	f := c.CelsiusToFahrenheit(0.00001)
	km := c.KilometersToMiles(238.2)
	fmt.Printf("Фаренгейт - %f. Километры - %f", f, km)
}