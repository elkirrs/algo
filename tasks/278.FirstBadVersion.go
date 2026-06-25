package tasks

import (
	"fmt"
)

func FirstBadVersionRun() {
	fmt.Println("278. First Bad Version")
	fmt.Println("https://leetcode.com/problems/first-bad-version/description/")
	fmt.Println("")

	version := 5
	fmt.Println("Input: version =", version)
	result := firstBadVersion(version)
	fmt.Println("Output:", result)

}

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left, right := 1, n

    for left < right {
        mid := left + (right - left)/2
        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}

var badVersion = 4

func isBadVersion(version int) bool {
	return version >= badVersion
}