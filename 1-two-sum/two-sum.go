package leetcode

// Solution using HashMap
// The time complexity for this is O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if j, ok := numMap[diff]; ok {
			return []int{j, i}
		}
		numMap[nums[i]] = i
	}
	return nil
}

//First Brute force attempt
// The time complexity for this is O(n^2)
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}
