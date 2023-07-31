package leetcode

// 辅助函数，用于生成子集
func backtrack(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		// 将当前数字加入子集中
		path = append(path, nums[i])

		// 递归生成下一个位置的子集
		backtrack(nums, i+1, path, res)

		// 回溯，恢复当前数字状态
		path = path[:len(path)-1]
	}
}

// Subsets 子集函数
func Subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	backtrack(nums, 0, path, &res)

	return res
}
