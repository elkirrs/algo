package tasks

import "fmt"

func TwoSumRun() {
	fmt.Println("Task: 1. Two Sum")
	fmt.Println("Link: https://leetcode.com/problems/two-sum/description/")
	fmt.Println("")

	nums := []int{-10, -8, -2, 1, 2, 5, 6}
	target := 0
	fmt.Println("Nums:", nums)
	fmt.Println("Target:", target)
	result := twoSum(nums, target)
	fmt.Println("Result:", result)


	nums = []int{3, 24, 50, 79, 88, 150, 345}
	target = 27
	fmt.Println("Nums:", nums)
	fmt.Println("Target:", target)

	result = twoSum(nums, target)
	fmt.Println("Result:", result)

}

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        d := target - v
        if idx, ex := m[d]; ex {
            return []int{idx, i}
        }
        m[v] = i
    }

    return []int{}
}


// Задача: Найти два числа в отсортированном массиве, которые в сумме дают заданное значение.

// Объяснение:
// 1. Используем два указателя: left с начала массива и right с конца.
// 2. Суммируем значения на позициях left и right.
// 3. Если сумма равна целевому значению, возвращаем индексы.
// 4. Если сумма меньше целевого значения, сдвигаем left вправо.
// 5. Если сумма больше целевого значения, сдвигаем right влево.