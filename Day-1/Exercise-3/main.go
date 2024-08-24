package main

import "fmt"

func findMax(slice []int) int {
	var max int = slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func findMin(slice []int) int {
	var min int = slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

func calculateSum(slice []int) int {
	var sum int
	for _, value := range slice {
		sum += value
	}
	return sum
}

func calculateAverage(slice []int) float64 {
	var sum int
	for _, value := range slice {
		sum += value
	}
	return float64(sum) / float64(len(slice))
}

func main() {
	var LengthOfSlice int
	fmt.Print("Enter the length of the slice: ")
	fmt.Scanln(&LengthOfSlice)

	var slice = make([]int, LengthOfSlice)
	for i := 0; i < LengthOfSlice; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&slice[i])
	}

	fmt.Println("Max:", findMax(slice))
	fmt.Println("Min:", findMin(slice))
	fmt.Println("Sum:", calculateSum(slice))
	fmt.Println("Average:", calculateAverage(slice))
}