/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	a := make([]int, 0)
	for i := 0; i < len(nums); i ++ {
			for j:= i + 1; j < len(nums); j++ {
					if nums[i] + nums[j] == target {
							a := append(a, i, j)
							return a
					}
			}
	}
	return a
}
// @lc code=end

