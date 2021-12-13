package leetcode

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	res := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

//
