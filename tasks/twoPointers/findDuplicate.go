package twoPointers

//nums = [2,5,9,6,9,3,8,9,7,1]

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
