// Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại


package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	inputLength := len(input)

	if inputLength&1 == 1 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

}