package main

import (
	"fmt"

	"example.com/twosum/twosum"
)

func main() {
	nums := []int{2, 7, 11, 15}
	res := twosum.TwoSum(nums, 9)
	fmt.Println(res)
}
