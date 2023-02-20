package main

var (
	nums    = []int{2, 7, 11, 15}
	target  = 9
	hashmap = make(map[int]int, len(nums))
	c       int
)

func twoSum(nums []int, target int) []int {

	for i, v := range nums {

		c = target - v

		j, ok := hashmap[c]

		if ok && i != hashmap[c] {
			return []int{i, j}
		}

		hashmap[v] = i
	}

	return nil
}
