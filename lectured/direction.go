package main

import "fmt"

type Direction string

func GetDirection(angle float64) Direction{
	switch {
	case angle >= 315:
		return "North"
	case angle > 0 && angle < 45:
		return "North"
	case angle >= 45 && angle < 135:
		return "East"
	case angle > 135 && angle < 255:
		return "South"
	case angle > 255 && angle < 315:
		return "West"
	default:
		return "None"
	}
}


func main(){
	angle := 288
	fmt.Println(GetDirection(float64(angle)))
}