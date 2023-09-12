package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	temp := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return nums[len(nums)/2]
}
