// Viết chương trình nhập giải bài toán twosum : https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	result := make([]int, 2)
	for i, num := range nums {
		if val, ok := dict[target-num]; ok {
			result[0] = val
			result[1] = i
			break
		}
		dict[num] = i
	}
	return result
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}
