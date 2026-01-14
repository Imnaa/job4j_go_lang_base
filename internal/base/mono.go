package base

/*
	Mono

Монотонный порядок
nums[i] >= nums[i+1] или
nums[i] <= nums[i+1]
*/
func Mono(nums []int) bool {
	size := len(nums)

	if size == 0 {
		return false
	}
	if size == 1 {
		return true
	}

	isIncreasing := true
	isDecreasing := true

	for i := 0; i < size-1; i++ {
		if nums[i] < nums[i+1] {
			isDecreasing = false
		}

		if nums[i] > nums[i+1] {
			isIncreasing = false
		}

		if !isIncreasing && !isDecreasing {
			return false
		}
	}

	return isIncreasing || isDecreasing
}
