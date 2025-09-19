package converter


func CelsiusToFahrenheit(c float64) float64{
	F := (c * 9 / 5) + 32
	return  F
}


func KilometersToMiles(km float64) float64{
	Miles := km * 0.62
	return  Miles
}