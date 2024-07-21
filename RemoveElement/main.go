package main

import "fmt"

var (
	nums = make([]int, 0, 4)
)

func main() {

	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 2)
	nums = append(nums, 3)

	n := removeElement(nums, 3)

	fmt.Printf("%v\n", n)

}

func removeElement(nums []int, val int) int {

	var n = len(nums) - 1
	var i = 0

	for i = 0; i <= n; i++ {

		if nums[i] == val {

			for nums[n] == val && n > i {
				n--
			}

			if nums[n] == val && n == i {
				return len(nums[:n])
			} else {
				nums[i] = nums[n]
			}
			n--
		}
	}

	return len(nums[:i])
}
