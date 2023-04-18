package leetcode

func containsDuplicate(nums []int) bool {

	numsMap := make(map[int]bool)

	for _, num := range nums {
		if _, exists := numsMap[num]; exists {
			return numsMap[num]
		}
		numsMap[num] = true
	}
	return false
}

//The time complexity for this solution is O(n)
