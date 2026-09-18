package main

import "fmt"

func twoSum(nums []int, target int) [] int {
	summ := make(map[int]int)

	for currIdx, currNum := range nums {
		needNum := target - currNum

		preIdx, found := summ[needNum]

		if found {
			return []int{preIdx, currIdx}
		}
		summ[currNum] = currIdx
	}

	return nil
}

func main() {
	nums := []int{4, 2, 6, 7}
	target := 11

	res := twoSum(nums, target)

	fmt.Println("Result: ", res)
}