package main

import "fmt"

func calculateArea(width, height float64) float64 {
	return width * height
}

func calculatePerimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func main() {
	var width, height float64
	fmt.Print("Enter width: ")
	fmt.Scanln(&width)
	fmt.Print("Enter height: ")
	fmt.Scanln(&height)
	fmt.Println("Area: ", calculateArea(width, height))
	fmt.Println("Perimeter: ", calculatePerimeter(width, height))
}
