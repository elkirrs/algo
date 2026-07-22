package tasks

func SlidingWindowMaximumRun() {
	PrintTask(
		"239. Sliding Window Maximum",
		"https://leetcode.com/problems/sliding-window-maximum/description/",
	)

	PrintScriptTwo(maxSlidingWindow, []int{1, 3, -1, -3, 5, 3, 6, 7}, 3) // 3,3,5,5,6,7
	PrintScriptTwo(maxSlidingWindow, []int{1}, 1)                        // 1
	PrintScriptTwo(maxSlidingWindow, []int{1, -1}, 1)                    // 1, -1
	PrintScriptTwo(maxSlidingWindow, []int{1, 3, 1, 2, 0, 5}, 3)         // 3,3,2,5
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	out := []int{}
	queue := []int{} // saving indexes

	for i := 0; i < len(nums); i++ {

		// remove an element that moved outside the window
		if len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		// remove all element less than the current one from the end
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		// prepare window
		if i >= k-1 {
			out = append(out, nums[queue[0]])
		}
	}

	return out
}
