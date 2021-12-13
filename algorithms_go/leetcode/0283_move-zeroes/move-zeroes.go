package leetcode

/*
0和非0元素不断对调
*/
func moveZeroes2(nums []int) {
	lenght := len(nums)
	if lenght <= 1 {
		return
	}
	pos := 0
	for i := 0; i < lenght; i++ {
		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}
}

/*
遇到非0元素就填入，最后看计数器的值，补0
*/
func moveZeroes(nums []int) {
	lenght := len(nums)
	if lenght <= 1 {
		return
	}
	pos := 0
	for i := 0; i < lenght; i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	for j := pos; j < lenght; j++ {
		nums[j] = 0
	}
}
