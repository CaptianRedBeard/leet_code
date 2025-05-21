package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	return hashMap(nums, target)

}

func hashMap(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, number := range nums {
		index2, exists := numMap[target-number]
		if exists {
			return []int{index2, index}
		} else {
			numMap[number] = index
		}
	}
	return nil
}

func bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// main function to run the program
func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println(result)
}
