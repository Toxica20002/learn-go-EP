// Viết chương trình nhập một slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp


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

func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	left, right := 0, len(slice)-1
	pivot := len(slice) / 2
	slice[pivot], slice[right] = slice[right], slice[pivot]
	for i := range slice {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	quickSort(slice[:left])
	quickSort(slice[left+1:])
	return slice
}

func sortAscending(slice []int) []int {
	return quickSort(slice)
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
	fmt.Println("Sorted Ascending:", sortAscending(slice))
}