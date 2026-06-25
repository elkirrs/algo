### 6. 4Sum

Задача: Найти все уникальные четверки чисел в массиве, которые в сумме дают заданное значение.

Пример кода на Go:

```go
import "sort"

func fourSum(nums []int, target int) [][]int {
sort.Ints(nums) // Sort the array to enable two-pointer approach
result := [][]int{}

// Iterate through the array using i as the first number of the quadruplet
for i := 0; i < len(nums)-3; i++ {
// Skip duplicate values for the first number in the quadruplet
if i > 0 && nums[i] == nums[i-1] {
continue
}

// Iterate through the array using j as the second number of the quadruplet
for j := i + 1; j < len(nums)-2; j++ {
// Skip duplicate values for the second number in the quadruplet
if j > i+1 && nums[j] == nums[j-1] {
continue
}

// Two-pointer approach for the third and fourth numbers
left, right := j+1, len(nums)-1
for left < right {
sum := nums[i] + nums[j] + nums[left] + nums[right]
if sum == target {
// Found a valid quadruplet
result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

// Skip duplicate values for the third and fourth numbers
for left < right && nums[left] == nums[left+1] {
left++
}
for left < right && nums[right] == nums[right-1] {
right--
}
left++
right--
} else if sum < target {
left++ // We need a larger sum, so increment left
} else {
right-- // We need a smaller sum, so decrement right
}
}
}
}
return result
}
```

Объяснение:

1. Мы сортируем массив, чтобы избежать повторений.
2. Для каждого числа перебираем все возможные тройки с двумя указателями для поиска четверок.
3. Если сумма равна целевому значению, добавляем комбинацию в результат и продолжаем сдвигать указатели, чтобы избежать
   повторений.

### 7. Sort Colors

Задача: Отсортировать массив, содержащий только 0, 1 и 2, без использования дополнительной памяти.

Пример кода на Go:

```go
func sortColors(nums []int) {
low, mid, high := 0, 0, len(nums)-1 // Initialize three pointers

// Traverse the array with the mid pointer
for mid <= high {
switch nums[mid] {
case 0:
// If the current element is 0, swap with the low pointer and move both low and mid forward
nums[low], nums[mid] = nums[mid], nums[low]
low++
mid++
case 1:
// If the current element is 1, just move mid forward
mid++
case 2:
// If the current element is 2, swap with the high pointer and move high backward
nums[mid], nums[high] = nums[high], nums[mid]
high--
}
}
}
```

Объяснение:

1. Используем три указателя: low для 0, mid для 1, и high для 2.
2. Перебираем элементы, сдвигая указатели в зависимости от значения текущего элемента.

### 8. Move Zeroes

Задача: Переместить все нули в конец массива, сохраняя порядок остальных элементов.

Пример кода на Go:

```go
func moveZeroes(nums []int) {
j := 0 // Pointer for the position of the next non-zero element
for i := 0; i < len(nums); i++ {
if nums[i] != 0 {
// Swap the non-zero element with the element at position j
nums[j], nums[i] = nums[i], nums[j]
j++ // Move the j pointer to the next position
}
}
}
```

Объяснение:

1. Используем один указатель j для отслеживания позиции для непустых элементов.
2. Когда находим ненулевой элемент, меняем его местами с элементом на позиции j и увеличиваем j.

### 9. Partition Labels

Задача: Разделить строку на подстроки, где каждая буква встречается только в одной подстроке.

Пример кода на Go:

```go
func partitionLabels(s string) []int {
// Step 1: Create a map to store the last occurrence index of each character
last := make(map[byte]int)
for i := 0; i < len(s); i++ {
last[s[i]] = i
}

// Step 2: Initialize variables to track the partitions
result := []int{}
start, end := 0, 0

// Step 3: Traverse the string to determine partitions
for i := 0; i < len(s); i++ {
// Update the end to be the farthest last occurrence of the current character
end = max(end, last[s[i]])

// If we've reached the end of a partition, calculate its length
if i == end {
result = append(result, i-start+1)
start = i + 1
}
}

return result
}

// Helper function to get the maximum of two numbers
func max(a, b int) int {
if a > b {
return a
}
return b
}
```

Объяснение:

1. Мы строим карту для хранения последнего вхождения каждого символа.
2. Используем два указателя для отслеживания диапазона подстроки, которая не содержит повторяющихся символов.
3. Когда мы достигаем последнего индекса для текущей подстроки, добавляем её длину в результат.

Каждый из этих примеров использует алгоритм с двумя указателями для эффективного решения задач с массивами и строками.
Это помогает значительно улучшить производительность по сравнению с наивными подходами.

# Cтроки

Алгоритмы для работы со строками используются для решения множества задач, таких как конвертация чисел в различные
системы счисления, проверка валидности строк или изменение их порядка. Рассмотрим несколько популярных задач со
строками, в которых применяются алгоритмы с разной сложностью и подходами, включая использование двух указателей,
хеш-таблиц и других техник.

### 1. Roman to Integer

Задача: Преобразовать римское число в целое число.

Римские цифры имеют специальные правила: числа могут быть написаны с использованием комбинации символов, где символы
перед более большими символами означают вычитание (например, IV = 4).

Пример кода на Go:

func romanToInt(s string) int {
roman := map[byte]int{
'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}
result := 0
for i := 0; i < len(s)-1; i++ {
if roman[s[i]] < roman[s[i+1]] {
result -= roman[s[i]] // Вычитание, если текущий символ меньше следующего
} else {
result += roman[s[i]]
}
}
result += roman[s[len(s)-1]] // Добавляем последний символ
return result
}

Объяснение:

1. Сначала создаем карту для римских цифр.
2. Проходим по строке, сравнивая текущую цифру с последующей.
3. Если текущая цифра меньше следующей, то из суммы вычитаем её (например, в случае IV, V - I).
4. Если текущая цифра больше или равна следующей, прибавляем её значение к результату.
5. Добавляем значение последнего символа в конце.

### 2. Integer to Roman

Задача: Преобразовать целое число в римское.

Римские цифры используют комбинации символов для представления чисел (например, 3999 = “MMMCMXCIX”).

Пример кода на Go:

func intToRoman(num int) string {
values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
result := ""
for i := 0; num > 0; i++ {
for num >= values[i] {
result += symbols[i]
num -= values[i]
}
}
return result
}

Объяснение:

1. Мы создаем два массива: один для значений римских цифр, другой — для их соответствующих символов.
2. Используем цикл, чтобы уменьшать число, начиная с самых крупных значений.
3. Когда текущее число больше или равно значению, мы добавляем символ в результат и вычитаем это значение из числа.
4. Повторяем процесс до тех пор, пока число не станет равным нулю.

### 3. Valid Palindrome

Задача: Проверить, является ли строка палиндромом (т.е. она читается одинаково слева направо и справа налево).

Пример кода на Go:

func isPalindrome(s string) bool {
left, right := 0, len(s)-1
for left < right {
if !isAlphaNumeric(s[left]) {
left++
} else if !isAlphaNumeric(s[right]) {
right--
} else if toLower(s[left]) != toLower(s[right]) {
return false
} else {
left++
right--
}
}
return true
}

func isAlphaNumeric(c byte) bool {
return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func toLower(c byte) byte {
if c >= 'A' && c <= 'Z' {
return c + 32
}
return c
}

Объяснение:

1. Используем два указателя: один с начала строки (left), второй — с конца (right).
2. Игнорируем все неалфавитные и нецифровые символы.
3. Сравниваем символы на left и right, если они не совпадают, возвращаем false.
4. Если символы совпадают, двигаем указатели в центр и продолжаем сравнение.

### 4. Valid Anagram

Задача: Проверить, являются ли две строки анаграммами друг друга (то есть содержат одинаковые символы, но в разном
порядке).

Пример кода на Go:

func isAnagram(s string, t string) bool {
if len(s) != len(t) {
return false
}
count := make(map[rune]int)
for _, c := range s {
count[c]++
}
for _, c := range t {
if count[c] == 0 {
return false
}
count[c]--
}
return true
}

Объяснение:

1. Если длины строк не совпадают, сразу возвращаем false.
2. Создаем карту для подсчета частоты символов в первой строке.
3. Для второй строки проверяем, существует ли такой символ в карте, и уменьшаем его счетчик.
4. Если все элементы совпали по количеству, возвращаем true.

### 5. Reverse String

Задача: Перевернуть строку.

Пример кода на Go:

func reverseString(s string) string {
chars := []rune(s) // Преобразуем строку в срез рунических значений для поддержки всех символов
left, right := 0, len(chars)-1
for left < right {
chars[left], chars[right] = chars[right], chars[left]
left++
right--
}
return string(chars)
}

Объяснение:

1. Преобразуем строку в срез рун для правильной обработки многобайтовых символов (например, кириллицы или эмодзи).
2. Используем два указателя, чтобы обменивать элементы с обоих концов строки.
3. Когда указатели встречаются, строка перевернута.

### 6. Zigzag Conversion

Задача: Преобразовать строку в зигзагообразный формат, прочитанный по строкам.

Пример кода на Go:

func convert(s string, numRows int) string {
if numRows == 1 {
return s
}
result := make([]string, numRows)
row, down := 0, false
for i := 0; i < len(s); i++ {
result[row] += string(s[i])
if row == 0 || row == numRows-1 {
down = !down
}
if down {
row++
} else {
row--
}
}
return strings.Join(result, "")
}

Объяснение:

1. Если строка должна быть только в одной строке, возвращаем её как есть.
2. Создаем массив строк, где каждая строка будет содержать символы для одной строки зигзагообразного порядка.
3. Используем два указателя для перемещения по строкам зигзагообразно: вверх и вниз.
4. Собираем строки обратно в одну с помощью strings.Join.

### 7. Length of Last Word

Задача: Найти длину последнего слова в строке, где слова разделяются пробелами.

Пример кода на Go:

func lengthOfLastWord(s string) int {
length := 0
for i := len(s) - 1; i >= 0; i-- {
if s[i] == ' ' && length > 0 {
break
} else if s[i] != ' ' {
length++
}
}
return length
}

Объяснение:

1. Идем с конца строки и начинаем считать длину последнего слова.
2. Если встречаем пробел и уже начали подсчитывать символы (то есть длина больше 0), завершаем поиск.
3. Возвращаем длину последнего слова.

Каждый из этих примеров использует различные методы для решения задач со строками. Алгоритмы могут варьироваться в
зависимости от задачи, но использование таких техник как два указателя, хеш-таблицы, и манипуляции со строками позволяют
эффективно решать эти задачи.

# Cвязные списки

Связные списки — это структуры данных, состоящие из узлов, где каждый узел содержит данные и указатель на следующий
узел. Они широко используются в различных задачах, таких как поиск, сортировка, удаление и манипуляции с элементами.
Рассмотрим несколько задач, связанных с обработкой связных списков, и алгоритмы, применяемые для их решения.

### 1. Middle of the Linked List

Задача: Найти середину связного списка.

Алгоритм: Для нахождения середины списка можно использовать два указателя: один будет двигаться на один шаг, а второй —
на два шага. Когда второй указатель достигнет конца, первый указатель будет находиться в середине.

Пример кода на Go:

type ListNode struct {
Val int
Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
slow, fast := head, head
for fast != nil && fast.Next != nil {
slow = slow.Next
fast = fast.Next.Next
}
return slow
}

Объяснение:

1. Используем два указателя: slow и fast.
2. Указатель fast двигается на два шага, а указатель slow на один.
3. Когда fast достигает конца списка, slow будет находиться в середине.

### 2. Linked List Cycle

Задача: Проверить, содержит ли связный список цикл.

Алгоритм: Мы можем использовать метод “черепахи и зайца”, где два указателя (один быстрый, другой медленный) движутся по
списку. Если список содержит цикл, два указателя встретятся.

Пример кода на Go:

func hasCycle(head *ListNode) bool {
slow, fast := head, head
for fast != nil && fast.Next != nil {
slow = slow.Next
fast = fast.Next.Next
if slow == fast {
return true
}
}
return false
}

Объяснение:

1. Два указателя: slow (медленный) и fast (быстрый) двигаются по списку.
2. Если fast догоняет slow, значит, есть цикл.
3. Если один из указателей достигает конца списка, цикла нет.

### 3. Remove Duplicates from Sorted List

Задача: Удалить дубликаты из отсортированного связного списка.

Алгоритм: Поскольку список отсортирован, все дубликаты будут расположены рядом. Проходим по списку и удаляем все
последующие одинаковые элементы.

Пример кода на Go:

func deleteDuplicates(head *ListNode) *ListNode {
current := head
for current != nil && current.Next != nil {
if current.Val == current.Next.Val {
current.Next = current.Next.Next
} else {
current = current.Next
}
}
return head
}

Объяснение:

1. Проходим по списку, сравнивая текущий элемент с следующим.
2. Если они одинаковы, удаляем следующий элемент.
3. Если не одинаковы, просто переходим к следующему элементу.

### 4. Reverse Linked List

Задача: Перевернуть связный список.

Алгоритм: Используем три указателя для переворачивания списка: prev, current и next. Идем по списку, меняем указатели на
каждый элемент.

Пример кода на Go:

func reverseList(head *ListNode) *ListNode {
var prev *ListNode
current := head
for current != nil {
next := current.Next
current.Next = prev
prev = current
current = next
}
return prev
}

Объяснение:

1. Используем три указателя: prev, current, и next.
2. Для каждого элемента меняем указатель на предыдущий элемент.
3. После завершения цикла prev будет указывать на новый “голову” списка.

### 5. Palindrome Linked List

Задача: Проверить, является ли связный список палиндромом.

Алгоритм: Разделить список пополам, развернуть вторую половину, а затем сравнить обе половины.

Пример кода на Go:

func isPalindrome(head *ListNode) bool {
if head == nil || head.Next == nil {
return true
}
// Найти середину списка
slow, fast := head, head
for fast != nil && fast.Next != nil {
slow = slow.Next
fast = fast.Next.Next
}
// Развернуть вторую половину списка
var prev *ListNode
for slow != nil {
next := slow.Next
slow.Next = prev
prev = slow
slow = next
}
// Сравнить две половины
left, right := head, prev
while right != nil {
if left.Val != right.Val {
return false
}
left = left.Next
right = right.Next
}
return true
}

Объяснение:

1. Находим середину списка с помощью двух указателей.
2. Разворачиваем вторую половину списка.
3. Сравниваем первую половину с развернутой второй половиной.

### 6. Merge Two Sorted Lists

Задача: Слить два отсортированных связных списка в один отсортированный.

Алгоритм: Используем два указателя, поочередно выбирая меньший элемент из двух списков.

Пример кода на Go:

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
dummy := &ListNode{}
current := dummy
for l1 != nil && l2 != nil {
if l1.Val < l2.Val {
current.Next = l1
l1 = l1.Next
} else {
current.Next = l2
l2 = l2.Next
}
current = current.Next
}
if l1 != nil {
current.Next = l1
}
if l2 != nil {
current.Next = l2
}
return dummy.Next
}

Объяснение:

1. Используем вспомогательную переменную dummy для построения нового списка.
2. Пошагово сравниваем элементы двух списков и прикрепляем меньший элемент к результату.
3. В конце присоединяем оставшиеся элементы, если они есть.

### 7. Merge K Sorted Lists

Задача: Слить k отсортированных связных списков в один отсортированный.

Алгоритм: Используем минимальную кучу (priority queue) для эффективного слияния нескольких списков.

Пример кода на Go:

import "container/heap"

type ListNode struct {
Val int
Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} { old := *h; n := len(old); x := old[n-1]; *h = old[0 : n-1]; return x }

func mergeKLists(lists []*ListNode) *ListNode {
minHeap := &MinHeap{}
heap.Init(minHeap)

    for _, list := range lists {
        if list != nil {
            heap.Push(minHeap, list)
        }
    }

    dummy := &ListNode{}
    current := dummy
    for minHeap.Len() > 0 {
        node := heap.Pop(minHeap).(*ListNode)
        current.Next = node
        current = current.Next
        if node.Next != nil {
            heap.Push(minHeap, node.Next)
        }
    }

    return dummy.Next

}

Объяснение:

1. Используем кучу (min-heap) для хранения узлов из каждого списка.
2. Пошагово извлекаем наименьший элемент из кучи и добавляем его в новый список.
3. После извлечения узла добавляем следующий элемент из того же списка, если он существует.

### 8. Delete Node in a Linked List

Задача: Удалить узел в связном списке, не зная о предыдущем узле.

Алгоритм: Скопировать данные следующего узла в текущий узел и удалить следующий узел.

Пример кода на Go:

func deleteNode(node *ListNode) {
node.Val = node.Next.Val
node.Next = node.Next.Next
}

Объяснение:

1. Копируем значение следующего узла в текущий узел.
2. Удаляем следующий узел, чтобы “переписать” текущий узел с его значением.

### 9. Add Two Numbers

Задача: Сложить два числа, представленных связными списками, где каждый узел представляет одну цифру.

Алгоритм: Идем по обоим спискам, суммируем соответствующие цифры и учитываем перенос.

Пример кода на Go:

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
dummy := &ListNode{}
current, carry := dummy, 0
for l1 != nil || l2 != nil || carry > 0 {
sum := carry
if l1 != nil {
sum += l1.Val
l1 = l1.Next
}
if l2 != nil {
sum += l2.Val
l2 = l2.Next
}
carry = sum / 10
current.Next = &ListNode{Val: sum % 10}
current = current.Next
}
return dummy.Next
}

Объяснение:

1. Используем переменную carry для учета переноса.
2. Суммируем соответствующие цифры двух списков.
3. Если есть перенос, добавляем его в следующую итерацию.

### 10. Sort List

Задача: Отсортировать связный список.

Алгоритм: Используем алгоритм сортировки слиянием, который подходит для связных списков.

Пример кода на Go:

func sortList(head *ListNode) *ListNode {
if head == nil || head.Next == nil {
return head
}
mid := getMiddle(head)
left := sortList(head)
right := sortList(mid)
return merge(left, right)
}

func getMiddle(head *ListNode) *ListNode {
slow, fast := head, head
for fast != nil && fast.Next != nil {
slow = slow.Next
fast = fast.Next.Next
}
return slow
}

func merge(left, right *ListNode) *ListNode {
dummy := &ListNode{}
current := dummy
for left != nil && right != nil {
if left.Val < right.Val {
current.Next = left
left = left.Next
} else {
current.Next = right
right = right.Next
}
current = current.Next
}
if left != nil {
current.Next = left
}
if right != nil {
current.Next = right
}
return dummy.Next
}

Объяснение:

1. Используем сортировку слиянием: делим список пополам, сортируем каждую часть, а затем сливаем.
2. Для разделения списка используем два указателя, чтобы найти середину.

Эти алгоритмы решают типичные задачи, связанные с манипуляциями с связанными списками.

# Деревья

Деревья — это важная структура данных, которая широко используется в алгоритмах и программировании. В бинарных деревьях
каждый узел имеет максимум два дочерних элемента (левый и правый). Алгоритмы, описанные ниже, связаны с различными
задачами, которые часто встречаются при работе с бинарными деревьями. Для каждого алгоритма приведен пример кода на Go и
его объяснение.

### 1. Binary Tree Inorder Traversal

Задача: Произвести обход бинарного дерева в порядке “левая вершина, корень, правая вершина” (inorder).

Алгоритм: Обходим дерево рекурсивно: сначала левое поддерево, затем корень, затем правое поддерево.

Пример кода на Go:

type TreeNode struct {
Val int
Left  *TreeNode
Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
var result []int
var helper func(*TreeNode)
helper = func(node *TreeNode) {
if node == nil {
return
}
helper(node.Left)
result = append(result, node.Val)
helper(node.Right)
}
helper(root)
return result
}

Объяснение:

- Мы используем рекурсивную функцию helper, чтобы пройти сначала по левому поддереву, затем добавить значение текущего
  узла в результат, и затем пройти по правому поддереву.

### 2. Binary Tree Preorder Traversal

Задача: Произвести обход бинарного дерева в порядке “корень, левая вершина, правая вершина” (preorder).

Алгоритм: Обходим дерево рекурсивно: сначала корень, затем левое поддерево, затем правое поддерево.

Пример кода на Go:

func preorderTraversal(root *TreeNode) []int {
var result []int
var helper func(*TreeNode)
helper = func(node *TreeNode) {
if node == nil {
return
}
result = append(result, node.Val)
helper(node.Left)
helper(node.Right)
}
helper(root)
return result
}

Объяснение:

- В этом обходе сначала записываем значение текущего узла в результат, затем рекурсивно обходим левое и правое
  поддеревья.

### 3. Binary Tree Postorder Traversal

Задача: Произвести обход бинарного дерева в порядке “левая вершина, правая вершина, корень” (postorder).

Алгоритм: Обходим дерево рекурсивно: сначала левое поддерево, затем правое поддерево, затем корень.

Пример кода на Go:

func postorderTraversal(root *TreeNode) []int {
var result []int
var helper func(*TreeNode)
helper = func(node *TreeNode) {
if node == nil {
return
}
helper(node.Left)
helper(node.Right)
result = append(result, node.Val)
}
helper(root)
return result
}

Объяснение:

- В этом обходе сначала обходим левое поддерево, затем правое поддерево, а затем добавляем значение узла в результат.

### 4. Binary Tree Level Order Traversal

Задача: Произвести обход бинарного дерева по уровням (уровневый обход).

Алгоритм: Используем очередь (FIFO) для уровня каждого узла. Сначала обрабатываем корень, затем для каждого узла
обрабатываем его детей.

Пример кода на Go:

import "container/list"

func levelOrder(root *TreeNode) [][]int {
if root == nil {
return nil
}

    result := [][]int{}
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        level := []int{}
        levelSize := queue.Len()

        for i := 0; i < levelSize; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            level = append(level, node.Val)

            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }

        result = append(result, level)
    }

    return result

}

Объяснение:

- Мы используем очередь для обработки узлов на каждом уровне. В каждой итерации вытягиваем все узлы на текущем уровне и
  добавляем их детей в очередь для обработки на следующем уровне.

### 5. Maximum Depth of Binary Tree

Задача: Найти максимальную глубину бинарного дерева.

Алгоритм: Используем рекурсию для нахождения максимальной глубины каждого поддерева.

Пример кода на Go:

func maxDepth(root *TreeNode) int {
if root == nil {
return 0
}
leftDepth := maxDepth(root.Left)
rightDepth := maxDepth(root.Right)
return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
if a > b {
return a
}
return b
}

Объяснение:

- Глубина дерева определяется как 1 (для текущего узла) плюс максимальная глубина левого и правого поддеревьев.

### 6. Minimum Depth of Binary Tree

Задача: Найти минимальную глубину бинарного дерева.

Алгоритм: Рекурсивно находим минимальную глубину левого и правого поддеревьев. Если одно из поддеревьев отсутствует,
глубина считается только по тому поддереву, которое присутствует.

Пример кода на Go:

func minDepth(root *TreeNode) int {
if root == nil {
return 0
}
if root.Left == nil {
return 1 + minDepth(root.Right)
}
if root.Right == nil {
return 1 + minDepth(root.Left)
}
leftDepth := minDepth(root.Left)
rightDepth := minDepth(root.Right)
return 1 + min(leftDepth, rightDepth)
}

func min(a, b int) int {
if a < b {
return a
}
return b
}

Объяснение:

- Минимальная глубина — это 1 плюс минимальная глубина левого и правого поддеревьев.

### 7. Path Sum

Задача: Проверить, существует ли путь, сумма значений которого равна заданной цели.

Алгоритм: Используем рекурсию для обхода дерева, вычитая значение каждого узла из оставшейся цели.

Пример кода на Go:

func hasPathSum(root *TreeNode, targetSum int) bool {
if root == nil {
return false
}
if root.Left == nil && root.Right == nil {
return root.Val == targetSum
}
targetSum -= root.Val
return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

Объяснение:

- Мы рекурсивно идем по дереву и вычитаем значение каждого узла из цели. Если достигнут листовой узел и оставшаяся цель
  равна 0, возвращаем true.

### 8. Symmetric Tree

Задача: Проверить, является ли бинарное дерево симметричным относительно своего центра.

Алгоритм: Используем рекурсию для проверки симметричности левого и правого поддеревьев.

Пример кода на Go:

func isSymmetric(root *TreeNode) bool {
if root == nil {
return true
}
return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
if t1 == nil && t2 == nil {
return true
}
if t1 == nil || t2 == nil {
return false
}
return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

Объяснение:

- Для проверки симметричности мы рекурсивно сравниваем левое и правое поддеревья, проверяя их зеркальность.

### 9. Same Tree

Задача: Проверить, являются ли два бинарных дерева одинаковыми.

Алгоритм: Рекурсивно сравниваем соответствующие узлы обоих деревьев.

Пример кода на Go:

func isSameTree(p *TreeNode, q *TreeNode) bool {
if p == nil && q == nil {
return true
}
if p == nil || q == nil {
return false
}
return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

Объяснение:

- Мы рекурсивно проверяем, равны ли значения текущих узлов и одинаковы ли их левые и правые поддеревья.

### 10. Binary Tree Paths

Задача: Найти все пути от корня до листьев.

Алгоритм: Рекурсивно строим пути, проходя по дереву.

Пример кода на Go:

func binaryTreePaths(root *TreeNode) []string {
var result []string
var helper func(*TreeNode, string)
helper = func(node *TreeNode, path string) {
if node == nil {
return
}
path += strconv.Itoa(node.Val)
if node.Left == nil && node.Right == nil {
result = append(result, path)
} else {
path += "->"
helper(node.Left, path)
helper(node.Right, path)
}
}
helper(root, "")
return result
}

Объяснение:

- Рекурсивно строим путь, добавляя значение каждого узла, пока не достигнем листа.

### 11. Validate Binary Search Tree

Задача: Проверить, является ли дерево бинарным деревом поиска.

Алгоритм: Используем рекурсивный обход с проверкой значений узлов относительно их родительских узлов.

Пример кода на Go:

func isValidBST(root *TreeNode) bool {
return isValidBSTHelper(root, math.MinInt, math.MaxInt)
}

func isValidBSTHelper(root *TreeNode, min, max int) bool {
if root == nil {
return true
}
if root.Val <= min || root.Val >= max {
return false
}
return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}

Объяснение:

- Мы рекурсивно проверяем, что значение каждого узла находится в допустимом диапазоне, соответствующем бинарному дереву
  поиска.

### 12. Lowest Common Ancestor of a Binary Tree

Задача: Найти наименьшего общего предка для двух узлов в бинарном дереве.

Алгоритм: Используем рекурсию для нахождения первого общего предка двух узлов.

Пример кода на Go:

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
if root == nil || root == p || root == q {
return root
}
left := lowestCommonAncestor(root.Left, p, q)
right := lowestCommonAncestor(root.Right, p, q)
if left == nil {
return right
}
if right == nil {
return left
}
return root
}

Объяснение:

- Мы рекурсивно ищем общих предков. Если один из поддеревьев возвращает ненулевое значение, это означает, что узлы
  находятся в одном поддереве.

### 13. Binary Tree Right Side View

Задача: Найти вид бинарного дерева с правой стороны.

Алгоритм: Используем уровень очереди для обхода дерева и записываем значения правых узлов на каждом уровне.

Пример кода на Go:

import "container/list"

func rightSideView(root *TreeNode) []int {
if root == nil {
return nil
}

    result := []int{}
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        levelSize := queue.Len()
        var rightMostNode int
        for i := 0; i < levelSize; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            rightMostNode = node.Val
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        result = append(result, rightMostNode)
    }

    return result

}

Объяснение:

- Мы используем уровень очереди для обработки каждого уровня дерева и записываем только правые узлы.

Все эти задачи являются основными примерами задач с бинарными деревьями.

# Хэш таблицы

Хэш-таблицы (или хэш-карты) — это структура данных, которая позволяет эффективно искать, вставлять и удалять элементы,
используя ключи. Хэш-таблицы широко применяются в различных задачах, где важна быстрая обработка данных.

Приведу примеры алгоритмов, которые используют хэш-таблицы, с подробным объяснением и примерами на языке Go.

### 1. Intersection of Two Arrays

Задача: Найти пересечение двух массивов.

Алгоритм: Используем хэш-таблицу, чтобы отслеживать элементы первого массива, а затем проверяем, встречаются ли они во
втором массиве.

Пример кода на Go:

func intersection(nums1 []int, nums2 []int) []int {
set := make(map[int]struct{})
for _, num := range nums1 {
set[num] = struct{}{}
}

    result := []int{}
    for _, num := range nums2 {
        if _, found := set[num]; found {
            result = append(result, num)
            delete(set, num)  // чтобы избежать повторений
        }
    }

    return result

}

Объяснение:

- Мы используем хэш-таблицу (set) для хранения элементов первого массива.
- Затем проверяем элементы второго массива и добавляем их в результат, если они есть в хэш-таблице. Удаляем элемент из
  хэш-таблицы, чтобы избежать повторений.

### 2. Two Sum

Задача: Найти два числа в массиве, которые в сумме дают заданное число.

Алгоритм: Используем хэш-таблицу, чтобы отслеживать элементы, которые мы уже встретили, и проверяем, можно ли найти
пару, которая дает нужную сумму.

Пример кода на Go:

func twoSum(nums []int, target int) []int {
mapNums := make(map[int]int)
for i, num := range nums {
complement := target - num
if idx, found := mapNums[complement]; found {
return []int{idx, i}
}
mapNums[num] = i
}
return nil
}

Объяснение:

- Мы проходим по массиву, и для каждого элемента вычисляем его “комплементарное” число, которое вместе с текущим
  элементом должно дать заданную сумму.
- Если комплемент уже есть в хэш-таблице, это значит, что мы нашли пару чисел, чья сумма равна заданной, и возвращаем
  индексы этих чисел.

### 3. Isomorphic Strings

Задача: Проверить, являются ли две строки изоморфными (то есть, у каждой буквы в первой строке есть уникальная замена в
другой строке).

Алгоритм: Используем две хэш-таблицы, чтобы отслеживать соответствия между символами двух строк.

Пример кода на Go:

func isIsomorphic(s string, t string) bool {
if len(s) != len(t) {
return false
}

    mapS := make(map[rune]rune)
    mapT := make(map[rune]rune)
    
    for i := 0; i < len(s); i++ {
        if mapS[rune(s[i])] != mapT[rune(t[i])] {
            return false
        }
        mapS[rune(s[i])] = rune(t[i])
        mapT[rune(t[i])] = rune(s[i])
    }
    
    return true

}

Объяснение:

- Для каждой пары символов в строках s и t мы проверяем, имеют ли они одинаковые соответствия в обеих хэш-таблицах.
- Если на каком-то шаге соответствие нарушается, возвращаем false.

### 4. Word Pattern

Задача: Проверить, соответствует ли строка шаблону.

Алгоритм: Используем хэш-таблицу, чтобы отслеживать соответствия между символами шаблона и словами в строке.

Пример кода на Go:

func wordPattern(pattern string, s string) bool {
words := strings.Split(s, " ")
if len(pattern) != len(words) {
return false
}

    mapPattern := make(map[byte]string)
    mapWords := make(map[string]byte)
    
    for i := 0; i < len(pattern); i++ {
        if existingWord, found := mapPattern[pattern[i]]; found {
            if existingWord != words[i] {
                return false
            }
        } else {
            mapPattern[pattern[i]] = words[i]
        }

        if existingChar, found := mapWords[words[i]]; found {
            if existingChar != pattern[i] {
                return false
            }
        } else {
            mapWords[words[i]] = pattern[i]
        }
    }
    return true

}

Объяснение:

- Мы используем две хэш-таблицы:

1. Одну для сопоставления символов шаблона с словами.
2. Другую для сопоставления слов с символами шаблона.

- Если для любого символа или слова сопоставление нарушается, возвращаем false.

### 5. Valid Sudoku

Задача: Проверить, является ли решетка судоку валидной (проверка строк, столбцов и подрешеток 3x3).

Алгоритм: Используем хэш-таблицы для проверки уникальности чисел в каждой строке, каждом столбце и каждом блоке 3x3.

Пример кода на Go:

func isValidSudoku(board [][]byte) bool {
rows := make([]map[byte]bool, 9)
cols := make([]map[byte]bool, 9)
boxes := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            num := board[i][j]
            if num == '.' {
                continue
            }

            boxIdx := (i/3)*3 + j/3
            if rows[i][num] || cols[j][num] || boxes[boxIdx][num] {
                return false
            }
            rows[i][num] = true
            cols[j][num] = true
            boxes[boxIdx][num] = true
        }
    }

    return true

}

Объяснение:

- Мы используем три хэш-таблицы:

1. Для каждой строки.
2. Для каждого столбца.
3. Для каждого блока 3x3.

- Для каждого числа проверяем, встречалось ли оно ранее в строке, столбце или блоке.

### 6. Group Anagrams

Задача: Группировать слова, которые являются анаграммами.

Алгоритм: Используем хэш-таблицу, где ключом будет отсортированное слово, а значением — список анаграмм.

Пример кода на Go:

func groupAnagrams(strs []string) [][]string {
anagrams := make(map[string][]string)

    for _, str := range strs {
        sorted := sortString(str)
        anagrams[sorted] = append(anagrams[sorted], str)
    }

    result := [][]string{}
    for _, group := range anagrams {
        result = append(result, group)
    }
    return result

}

func sortString(s string) string {
chars := []rune(s)
sort.Slice(chars, func(i, j int) bool {
return chars[i] < chars[j]
})
return string(chars)
}

Объяснение:

- Для каждого слова мы создаем его отсортированную версию и используем её как ключ в хэш-таблице.
- Слова с одинаковыми отсортированными версиями попадают в одну группу.

### 7. LRU Cache

Задача: Реализовать кэш с алгоритмом “наименее недавно использованные” (LRU), который удаляет старые элементы, когда
память переполнена.

Алгоритм: Используем хэш-таблицу для быстрого поиска и связанный список для отслеживания порядка использования
элементов.

Пример кода на Go:

type LRUCache struct {
capacity int
cache map[int]*ListNode
order    *DoubleLinkedList
}

type ListNode struct {
key, value int
prev, next *ListNode
}

type DoubleLinkedList struct {
head, tail *ListNode
}

func Constructor(capacity int) LRUCache {
cache := make(map[int]*ListNode)
order := &DoubleLinkedList{head: &ListNode{}, tail: &ListNode{}}
order.head.next = order.tail
order.tail.prev = order.head
return LRUCache{capacity: capacity, cache: cache, order: order}
}

func (this *LRUCache) Get(key int) int {
if node, found := this.cache[key]; found {
this.moveToFront(node)
return node.value
}
return -1
}

func (this *LRUCache) Put(key int, value int) {
if node, found := this.cache[key]; found {
node.value = value
this.moveToFront(node)
} else {
if len(this.cache) == this.capacity {
this.removeLeastUsed()
}
node := &ListNode{key: key, value: value}
this.cache[key] = node
this.addToFront(node)
}
}

func (this *LRUCache) moveToFront(node *ListNode) {
node.prev.next = node.next
node.next.prev = node.prev
node.next = this.order.head.next
node.prev = this.order.head
this.order.head.next.prev = node
this.order.head.next = node
}

func (this *LRUCache) addToFront(node *ListNode) {
node.next = this.order.head.next
node.prev = this.order.head
this.order.head.next.prev = node
this.order.head.next = node
}

func (this *LRUCache) removeLeastUsed() {
node := this.order.tail.prev
this.cache[node.key] = nil
node.prev.next = this.order.tail
this.order.tail.prev = node.prev
}

Объяснение:

- Используем хэш-таблицу для хранения элементов кэша.
- Связанный список помогает отслеживать порядок использования элементов, а старые элементы удаляются, когда кэш
  переполняется.

### 8. Top K Frequent Elements

Задача: Найти K самых частых элементов в массиве.

Алгоритм: Используем хэш-таблицу для подсчета частоты элементов и сортируем их по частоте.

Пример кода на Go:

func topKFrequent(nums []int, k int) []int {
freq := make(map[int]int)
for _, num := range nums {
freq[num]++
}

    var pairs []struct {
        num  int
        count int
    }

    for num, count := range freq {
        pairs = append(pairs, struct {
            num  int
            count int
        }{num, count})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].count > pairs[j].count
    })

    result := []int{}
    for i := 0; i < k; i++ {
        result = append(result, pairs[i].num)
    }
    return result

}

Объяснение:

- Мы используем хэш-таблицу для подсчета частоты каждого элемента.
- Затем сортируем элементы по частоте и выбираем K самых популярных.

# Матрицы

Задачи с матрицами

Работа с матрицами — важная тема в программировании. Во многих задачах нужно обрабатывать элементы матрицы в
определенном порядке или выполнять операции, которые требуют знания о соседях или позициях элементов. Рассмотрим
несколько задач с матрицами и алгоритмы для их решения с примерами на Go.

### 1. Matrix Diagonal Sum

Задача: Найти сумму элементов на главной диагонали и побочной диагонали матрицы.

Алгоритм: Мы будем итеративно проходить по строкам матрицы и для каждой строки добавлять элементы из главной и побочной
диагонали.

Пример кода на Go:

func diagonalSum(mat [][]int) int {
n := len(mat)
sum := 0
for i := 0; i < n; i++ {
sum += mat[i][i]        // главная диагональ
sum += mat[i][n-1-i]    // побочная диагональ
}

    // Если размер матрицы нечетный, вычитаем центральный элемент (он был добавлен дважды)
    if n%2 == 1 {
        sum -= mat[n/2][n/2]
    }

    return sum

}

Объяснение:

- Мы проходим по матрице и суммируем элементы на главной и побочной диагоналях.
- Если размер матрицы нечетный, центральный элемент будет добавлен дважды, и его нужно вычесть.

### 2. Valid Sudoku

Задача: Проверить, является ли доска судоку валидной (проверка строк, столбцов и подрешеток 3x3).

Алгоритм: Используем хэш-таблицы для проверки уникальности чисел в каждой строке, столбце и подрешетке 3x3.

Пример кода на Go:

func isValidSudoku(board [][]byte) bool {
rows := make([]map[byte]bool, 9)
cols := make([]map[byte]bool, 9)
boxes := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            num := board[i][j]
            if num == '.' {
                continue
            }

            boxIdx := (i/3)*3 + j/3
            if rows[i][num] || cols[j][num] || boxes[boxIdx][num] {
                return false
            }
            rows[i][num] = true
            cols[j][num] = true
            boxes[boxIdx][num] = true
        }
    }

    return true

}

Объяснение:

- Мы используем три хэш-таблицы для отслеживания чисел в строках, столбцах и 3x3 блоках.
- Если повторяется число в одной из этих областей, возвращаем false.

### 3. 01 Matrix

Задача: Дано бинарное изображение, заполненное 0 и 1. Нужно найти для каждого элемента матрицы минимальное расстояние до
ближайшего 0.

Алгоритм: Мы будем использовать метод многократного обхода с двумя проходами: сначала слева направо и сверху вниз, затем
справа налево и снизу вверх. Это позволит эффективно вычислить минимальное расстояние.

Пример кода на Go:

func updateMatrix(mat [][]int) [][]int {
m, n := len(mat), len(mat[0])
dist := make([][]int, m)
for i := 0; i < m; i++ {
dist[i] = make([]int, n)
for j := 0; j < n; j++ {
if mat[i][j] == 1 {
dist[i][j] = math.MaxInt
} else {
dist[i][j] = 0
}
}
}

    // Первый проход: слева направо, сверху вниз
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i > 0 {
                dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
            }
            if j > 0 {
                dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
            }
        }
    }

    // Второй проход: справа налево, снизу вверх
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if i < m-1 {
                dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
            }
            if j < n-1 {
                dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
            }
        }
    }

    return dist

}

func min(a, b int) int {
if a < b {
return a
}
return b
}

Объяснение:

- Первый проход заполняет расстояния, двигаясь слева направо и сверху вниз.
- Второй проход исправляет расстояния, двигаясь справа налево и снизу вверх, что позволяет учесть минимальные расстояния
  от ближайших 0.

### 4. Maximal Square

Задача: Найти максимальный квадрат в бинарной матрице, состоящий только из 1.

Алгоритм: Мы будем использовать динамическое программирование для вычисления максимальной стороны квадрата на каждом
элементе матрицы.

Пример кода на Go:

func maximalSquare(matrix [][]byte) int {
if len(matrix) == 0 || len(matrix[0]) == 0 {
return 0
}

    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    maxSide := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
                }
                maxSide = max(maxSide, dp[i][j])
            }
        }
    }

    return maxSide * maxSide

}

func min(a, b int) int {
if a < b {
return a
}
return b
}

func max(a, b int) int {
if a > b {
return a
}
return b
}

Объяснение:

- Для каждого элемента в матрице мы находим максимальный квадрат, используя минимальное значение из соседних ячеек,
  включая саму текущую ячейку.
- dp[i][j] хранит размер максимального квадрата, который заканчивается в клетке (i, j).

### 5. Set Matrix Zeroes

Задача: Если в матрице есть хотя бы один элемент, равный 0, то обнуляем всю строку и столбец, в котором находится этот
элемент.

Алгоритм: Для этого можно использовать два массива: один для строк и один для столбцов, чтобы отслеживать, какие строки
и столбцы нужно обнулить.

Пример кода на Go:

func setZeroes(matrix [][]int) {
m, n := len(matrix), len(matrix[0])
rows := make([]bool, m)
cols := make([]bool, n)

    // Отметить строки и столбцы, которые нужно обнулить
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rows[i] = true
                cols[j] = true
            }
        }
    }

    // Обнулить строки
    for i := 0; i < m; i++ {
        if rows[i] {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }

    // Обнулить столбцы
    for j := 0; j < n; j++ {
        if cols[j] {
            for i := 0; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }

}

Объяснение:

- Мы используем два массива, чтобы отметить строки и столбцы, содержащие 0.
- Затем в два прохода обнуляем нужные строки и столбцы.

### 6. Spiral Matrix

Задача: Прочитать матрицу по спирали.

Алгоритм: Прочитаем элементы матрицы, начиная с верхнего левого угла и двигаясь по спирали.

Пример кода на Go:

func spiralOrder(matrix [][]int) []int {
result := []int{}
if len(matrix) == 0 || len(matrix[0]) == 0 {
return result
}

    top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
    for top <= bottom && left <= right {
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++

        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--

        if top <= bottom {
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
            bottom--
        }

        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }

    return result

}

Объяснение:

- Мы проходим по матрице, начиная с внешнего слоя и постепенно сужая область для обхода, пока не пройдем всю матрицу.

### 7. Rotate Image

Задача: Повернуть квадратную матрицу на 90 градусов по часовой стрелке.

Алгоритм: Сначала транспонируем матрицу, а затем инвертируем строки.

Пример кода на Go:

func rotate(matrix [][]int) {
n := len(matrix)
// Транспонируем матрицу
for i := 0; i < n; i++ {
for j := i + 1; j < n; j++ {
matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
}
}

    // Инвертируем строки
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }

}

Объяснение:

- Транспонируем матрицу (обменяем элементы по диагонали).
- Инвертируем строки, чтобы добиться поворота на 90 градусов.

Каждый из этих алгоритмов показывает, как эффективно работать с матрицами, используя различные подходы, включая
динамическое программирование, хэширование, и прямолинейные итерации по элементам.

# Очередь и стек

Очереди и Стек (Queue and Stack)

Очереди и стеки являются основными структурами данных, которые используются для решения множества задач в
программировании. Стек — это структура данных, которая работает по принципу LIFO (Last In, First Out), а очередь — по
принципу FIFO (First In, First Out). Рассмотрим несколько задач, где эти структуры данных играют важную роль, а также
приведем примеры решений на языке Go.

### 1. Valid Parentheses

Задача: Проверить, являются ли скобки в строке правильными (открывающая скобка должна иметь соответствующую
закрывающую).

Алгоритм: Используем стек, чтобы отслеживать открывающие скобки. Когда встречаем закрывающую скобку, проверяем,
соответствует ли она последней открывающей скобке в стеке.

Пример кода на Go:

func isValid(s string) bool {
stack := []rune{}
for _, ch := range s {
if ch == '(' || ch == '[' || ch == '{' {
stack = append(stack, ch)
} else {
if len(stack) == 0 {
return false
}
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]
if (ch == ')' && top != '(') || (ch == ']' && top != '[') || (ch == '}' && top != '{') {
return false
}
}
}
return len(stack) == 0
}

Объяснение:

- Каждый раз, когда встречаем открывающую скобку, добавляем ее в стек.
- Когда встречаем закрывающую скобку, проверяем, соответствует ли она последней открытой скобке (удаленной из стека).
- Если в конце стек пуст, значит скобки правильно сбалансированы.

### 2. Min Stack

Задача: Реализовать стек с поддержкой операции поиска минимального элемента за время O(1).

Алгоритм: В дополнение к обычному стеку поддерживаем второй стек, который будет хранить минимальные значения на каждом
шаге.

Пример кода на Go:

type MinStack struct {
stack []int
minStack []int
}

func Constructor() MinStack {
return MinStack{}
}

func (this *MinStack) Push(x int) {
this.stack = append(this.stack, x)
if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1] {
this.minStack = append(this.minStack, x)
}
}

func (this *MinStack) Pop() {
top := this.stack[len(this.stack)-1]
this.stack = this.stack[:len(this.stack)-1]
if top == this.minStack[len(this.minStack)-1] {
this.minStack = this.minStack[:len(this.minStack)-1]
}
}

func (this *MinStack) Top() int {
return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
return this.minStack[len(this.minStack)-1]
}

Объяснение:

- Для поиска минимального значения используем дополнительный стек minStack, в котором храним минимальные элементы.
- При добавлении элемента в основной стек проверяем, нужно ли добавить его в minStack.
- При удалении элемента из стека, если он является минимальным, удаляем его и из minStack.

### 3. Daily Temperatures

Задача: Для каждого дня в массиве температур определить, сколько дней нужно ждать, чтобы температура стала выше.

Алгоритм: Используем стек, чтобы отслеживать индексы температур, и для каждого дня вычисляем, сколько дней нужно
подождать до более высокой температуры.

Пример кода на Go:

func dailyTemperatures(T []int) []int {
result := make([]int, len(T))
stack := []int{}

    for i := 0; i < len(T); i++ {
        for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
            index := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[index] = i - index
        }
        stack = append(stack, i)
    }
    
    return result

}

Объяснение:

- Для каждого дня мы проверяем, есть ли в стеке более старые дни с температурой, меньшей чем текущая.
- Если такая температура найдена, то мы вычисляем, сколько дней потребуется, чтобы температура стала выше, и удаляем
  старые дни из стека.
- После этого добавляем текущий день в стек для дальнейших сравнений.

### 4. Implement Queue Using Stacks

Задача: Реализовать очередь, используя два стека.

Алгоритм: Для реализации очереди используем два стека: один для вставки элементов, а второй — для удаления. Элементы
перемещаются из одного стека в другой при необходимости.

Пример кода на Go:

type MyQueue struct {
stackIn  []int
stackOut []int
}

func Constructor() MyQueue {
return MyQueue{}
}

func (this *MyQueue) Push(x int) {
this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
if len(this.stackOut) == 0 {
for len(this.stackIn) > 0 {
this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
this.stackIn = this.stackIn[:len(this.stackIn)-1]
}
}
val := this.stackOut[len(this.stackOut)-1]
this.stackOut = this.stackOut[:len(this.stackOut)-1]
return val
}

func (this *MyQueue) Peek() int {
if len(this.stackOut) == 0 {
for len(this.stackIn) > 0 {
this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
this.stackIn = this.stackIn[:len(this.stackIn)-1]
}
}
return this.stackOut[len(this.stackOut)-1]
}

func (this *MyQueue) Empty() bool {
return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

Объяснение:

- Мы используем два стека для реализации очереди.
- Когда нужно удалить элемент, мы перемещаем все элементы из стека stackIn в stackOut (если stackOut пуст), после чего
  извлекаем верхний элемент из stackOut.
- Таким образом, элементы в stackOut всегда идут в том порядке, в каком должны быть извлечены.

### 5. Implement Stack Using Queue

Задача: Реализовать стек, используя две очереди.

Алгоритм: Используем две очереди, одна из которых будет использоваться для хранения элементов стека, а другая — для
временного хранения элементов при операциях.

Пример кода на Go:

type MyStack struct {
queue1, queue2 []int
}

func Constructor() MyStack {
return MyStack{}
}

func (this *MyStack) Push(x int) {
this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
for len(this.queue1) > 1 {
this.queue2 = append(this.queue2, this.queue1[0])
this.queue1 = this.queue1[1:]
}
val := this.queue1[0]
this.queue1 = this.queue1[1:]
this.queue1, this.queue2 = this.queue2, this.queue1
return val
}

func (this *MyStack) Top() int {
return this.queue1[len(this.queue1)-1]
}

func (this *MyStack) Empty() bool {
return len(this.queue1) == 0
}

Объяснение:

- Мы используем две очереди для симуляции работы стека.
- При добавлении элемента он просто добавляется в одну очередь.
- При удалении элемента, мы перемещаем все элементы, кроме последнего, во вторую очередь, а затем извлекаем последний
  элемент из первой очереди.

### 6. Minimum Remove to Make Valid Parentheses

Задача: Найти минимальное количество удалений, чтобы строка с скобками стала валидной.

Алгоритм: Используем стек для отслеживания открытых скобок, а затем находим все индексы, которые необходимо удалить.

Пример кода на Go:

func minRemoveToMakeValid(s string) string {
stack := []int{}
result := []byte(s)

    for i, ch := range s {
        if ch == '(' {
            stack = append(stack, i)
        } else if ch == ')' {
            if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
                stack = stack[:len(stack)-1]
            } else {
                result[i] = '#'
            }
        }
    }

    for len(stack) > 0 {
        result[stack[len(stack)-1]] = '#'
        stack = stack[:len(stack)-1]
    }

    res := []byte{}
    for _, ch := range result {
        if ch != '#' {
            res = append(res, ch)
        }
    }
    
    return string(res)

}

Объяснение:

- Мы используем стек, чтобы отслеживать индексы открывающих скобок.
- Когда находим закрывающую скобку, проверяем, можно ли ее закрыть, и если нет, помечаем ее для удаления.
- В конце удаляем все оставшиеся открытые скобки, которые не были закрыты.

# Битовые манипуляции

Битовые манипуляции (Bit Manipulation)

Битовые манипуляции — это операции, которые работают с отдельными битами чисел. Они часто используются в задачах,
связанных с оптимизацией, криптографией и низкоуровневым программированием. Рассмотрим несколько популярных задач и их
решение с помощью битовых операций на языке Go.

### 1. Power of Two

Задача: Проверить, является ли число степенью двойки.

Алгоритм: Число является степенью двойки, если оно больше нуля и при побитовой операции “И” с его предыдущим числом
результат равен нулю. Это можно записать как:
n > 0 && (n & (n - 1)) == 0.

Пример кода на Go:

func isPowerOfTwo(n int) bool {
return n > 0 && (n & (n - 1)) == 0
}

Объяснение:

- Если число является степенью двойки, то в его двоичном представлении только один бит равен 1. Например, 16 в двоичной
  системе — это 10000.
- При вычитании 1 из числа и выполнении побитовой операции “И”, все биты должны обнулиться, если число является степенью
  двойки.

### 2. Power of Four

Задача: Проверить, является ли число степенью четырех.

Алгоритм: Число является степенью четырех, если оно является степенью двух, и его двоичное представление имеет нечетное
количество нулей после первого бита. Можно использовать побитовую операцию, проверяя, что (n & (n - 1)) == 0 (чтобы
удостовериться, что это степень двойки), и n % 3 == 1 (чтобы проверить, что это степень четырех).

Пример кода на Go:

func isPowerOfFour(n int) bool {
return n > 0 && (n & (n - 1)) == 0 && n % 3 == 1
}

Объяснение:

- Если число является степенью 4, то оно одновременно является степенью 2, но также нужно, чтобы остаток от деления на 3
  равнялся 1. Это связано с тем, что степень 4 выглядит как 1, 4, 16, 64, и так далее в двоичной системе.

### 3. Counting Bits

Задача: Подсчитать количество единичных битов в числе.

Алгоритм: Используем побитовую операцию для удаления самого младшего бита 1 и подсчитываем, сколько раз это можно
сделать.

Пример кода на Go:

func hammingWeight(n uint32) int {
count := 0
for n > 0 {
count++
n = n & (n - 1)
}
return count
}

Объяснение:

- Операция n & (n - 1) удаляет самый младший единичный бит из числа. Мы повторяем эту операцию, пока все биты не станут
  равными нулю, и каждый раз увеличиваем счетчик.

### 4. Reverse Bits

Задача: Развернуть биты числа.

Алгоритм: Разворачиваем биты числа, начиная с самого младшего бита и двигаясь к старшему, и постепенно строим новое
число с перевернутыми битами.

Пример кода на Go:

func reverseBits(n uint32) uint32 {
var result uint32
for i := 0; i < 32; i++ {
result <<= 1
result |= n & 1
n >>= 1
}
return result
}

Объяснение:

- Мы поочередно перемещаем биты числа в новое место, начиная с младшего бита. С каждым циклом сдвигаем результат на одну
  позицию влево и добавляем младший бит числа.
- Таким образом, мы строим новое число с развернутыми битами.

### 5. Hamming Distance

Задача: Найти расстояние Хэмминга между двумя числами (сколько битов различаются в их двоичных представлениях).

Алгоритм: Используем операцию “И” для нахождения отличий между числами, а затем подсчитываем количество единичных битов
в результате.

Пример кода на Go:

func hammingDistance(x, y int) int {
diff := x ^ y
count := 0
for diff > 0 {
count++
diff &= diff - 1
}
return count
}

Объяснение:

- Операция x ^ y дает число, где единичные биты показывают различия между числами x и y. Затем мы подсчитываем
  количество единичных битов в этом числе с помощью операции diff &= diff - 1, которая удаляет самый младший единичный
  бит.

### 6. Total Hamming Distance

Задача: Найти общее расстояние Хэмминга между всеми парами чисел в массиве.

Алгоритм: Для каждого бита вычисляем, сколько чисел имеют единичный бит в этой позиции, и на основе этого подсчитываем
общее расстояние Хэмминга.

Пример кода на Go:

func totalHammingDistance(nums []int) int {
total := 0
for i := 0; i < 32; i++ {
count := 0
for _, num := range nums {
if (num >> i) & 1 == 1 {
count++
}
}
total += count * (len(nums) - count)
}
return total
}

Объяснение:

- Для каждого бита мы подсчитываем, сколько чисел имеют единичный бит в этой позиции. Общее количество различий для
  этого бита — это count * (len(nums) - count), где count — количество чисел с единичным битом в данной позиции.
- Мы суммируем эти значения для всех 32 битов, чтобы получить общее расстояние Хэмминга.

Итоги

В задачах с битовыми манипуляциями важно эффективно работать с двоичным представлением чисел. Битовые операции, такие
как побитовые сдвиги, операции “И”, “Или” и “Исключающее И”, позволяют оптимизировать решение многих задач, снижая время
работы алгоритмов и минимизируя использование дополнительной памяти.

# Cкользящие окна

Скользящие окна (Sliding Window)

Алгоритм “скользящего окна” используется для решения задач, где нужно обрабатывать подмножества или подстроки данных,
проходя по ним один за одним, с использованием фиксированного размера окна, которое “скользит” по данным. Этот подход
часто позволяет улучшить эффективность работы с массивами или строками.

### 1. Longest Substring Without Repeating Characters

Задача: Найти длину самой длинной подстроки, в которой все символы уникальны.

Алгоритм: Используем два указателя (левый и правый) для представления окна. Перемещаем правый указатель, добавляя
символы в окно, и если символ уже встречался в окне, перемещаем левый указатель, чтобы исключить повторяющиеся символы.

Пример кода на Go:

func lengthOfLongestSubstring(s string) int {
charMap := make(map[rune]int)
left := 0
maxLength := 0

    for right, char := range s {
        if idx, found := charMap[char]; found && idx >= left {
            left = idx + 1
        }
        charMap[char] = right
        maxLength = max(maxLength, right - left + 1)
    }

    return maxLength

}

func max(a, b int) int {
if a > b {
return a
}
return b
}

Объяснение:

- Мы используем map для отслеживания индексов символов в текущем окне.
- Когда находим повторяющийся символ, сдвигаем левый указатель вправо, исключая предыдущую позицию этого символа.
- Увеличиваем максимальную длину подстроки на каждом шаге, если текущая длина окна больше.

### 2. Longest Repeating Character Replacement

Задача: Дано строка, в которой нужно заменить не более k символов, чтобы получить самую длинную строку, в которой все
символы одинаковы.

Алгоритм: Используем два указателя для определения окна, в котором количество символов, отличных от самого частого, не
превышает k. Расширяем окно и сужаем его, если количество замен превышает k.

Пример кода на Go:

func characterReplacement(s string, k int) int {
count := make(map[byte]int)
left := 0
maxCount := 0
maxLength := 0

    for right := 0; right < len(s); right++ {
        count[s[right]]++
        maxCount = max(maxCount, count[s[right]])

        if (right - left + 1) - maxCount > k {
            count[s[left]]--
            left++
        }

        maxLength = max(maxLength, right - left + 1)
    }

    return maxLength

}

func max(a, b int) int {
if a > b {
return a
}
return b
}

Объяснение:

- Мы сохраняем количество каждого символа в окне и отслеживаем, какой символ встречается наиболее часто.
- Если разница между размером окна и максимальной частотой символов больше k, сдвигаем левый указатель вправо, чтобы
  уменьшить количество замен.
- Вычисляем максимальную длину окна с допустимыми заменами.

### 3. Fruit Into Baskets

Задача: У нас есть массив фруктов, представленных числами. Мы можем держать только два типа фруктов в корзине. Нужно
найти наибольшее количество фруктов, которые можно собрать с этой ограниченной корзиной.

Алгоритм: Мы используем два указателя для скользящего окна, чтобы отслеживать два типа фруктов и расширять окно, пока не
встретим больше двух типов фруктов.

Пример кода на Go:

func totalFruit(fruits []int) int {
fruitCount := make(map[int]int)
left := 0
maxFruits := 0

    for right := 0; right < len(fruits); right++ {
        fruitCount[fruits[right]]++

        for len(fruitCount) > 2 {
            fruitCount[fruits[left]]--
            if fruitCount[fruits[left]] == 0 {
                delete(fruitCount, fruits[left])
            }
            left++
        }

        maxFruits = max(maxFruits, right - left + 1)
    }

    return maxFruits

}

func max(a, b int) int {
if a > b {
return a
}
return b
}

Объяснение:

- Мы используем map для подсчета фруктов в текущем окне.
- Как только в окне появляется более двух типов фруктов, мы начинаем сдвигать левый указатель, чтобы уменьшить
  количество типов фруктов.
- На каждом шаге обновляем максимальное количество фруктов в корзине.

### 4. Find All Anagrams in a String

Задача: Найти все начала анAGRAM в строке, то есть найти все позиции, с которых начинаются подстроки, являющиеся
анаграммами другой строки.

Алгоритм: Мы используем два окна одинаковой длины — одно для строки, а другое для ее анаграммы. Сравниваем частотный
состав символов в этих окнах.

Пример кода на Go:

func findAnagrams(s string, p string) []int {
result := []int{}
if len(s) < len(p) {
return result
}

    pCount := make([]int, 26)
    sCount := make([]int, 26)

    for _, ch := range p {
        pCount[ch - 'a']++
    }

    for i := 0; i < len(s); i++ {
        sCount[s[i] - 'a']++
        
        if i >= len(p) {
            sCount[s[i - len(p)] - 'a']--
        }

        if equal(sCount, pCount) {
            result = append(result, i - len(p) + 1)
        }
    }

    return result

}

func equal(a, b []int) bool {
for i := 0; i < 26; i++ {
if a[i] != b[i] {
return false
}
}
return true
}

Объяснение:

- Для каждого символа в строке p мы обновляем частотный массив.
- Для строки s используем окно той же длины и обновляем частотный массив.
- Сравниваем два массива частот и добавляем позиции, где частоты совпадают.

### 5. Sliding Window Maximum

Задача: Найти максимальное значение в каждом окне фиксированного размера в массиве.

Алгоритм: Используем очередь (или двустороннюю очередь), чтобы хранить индексы элементов, подходящих для текущего окна.
Поддерживаем элементы в очереди в порядке убывания.

Пример кода на Go:

import "container/deque"

func maxSlidingWindow(nums []int, k int) []int {
var result []int
dq := deque.New()

    for i := 0; i < len(nums); i++ {
        // Убираем элементы, выходящие за пределы окна
        if dq.Len() > 0 && dq.PeekFront().(int) < i - k + 1 {
            dq.RemoveFirst()
        }

        // Убираем элементы, меньшие текущего, чтобы гарантировать, что в очереди будет только максимальное значение
        for dq.Len() > 0 && nums[dq.PeekLast().(int)] < nums[i] {
            dq.RemoveLast()
        }

        dq.Append(i)

        // Если окно полно, добавляем максимальное значение в результат
        if i >= k - 1 {
            result = append(result, nums[dq.PeekFront().(int)])
        }
    }

    return result

}

Объяснение:

- Мы используем очередь для хранения индексов элементов в окне, при этом всегда удаляем индексы старых элементов и
  элементы меньшие текущего.
- Таким образом, максимальный элемент всегда будет находиться в начале очереди.

Итоги

Скользящее окно — мощный подход для оптимизации задач, связанных с последовательностями данных. Он позволяет
обрабатывать большие массивы или строки за время O(n), что делает решение более эффективным по сравнению с наивными
методами, где проверяются все возможные подстроки или подмассивы.

# Поиск с возвратом

Поиск с возвратом (Backtracking)

Поиск с возвратом — это метод решения задач, где мы пытаемся все возможные варианты решения и “возвращаемся” (откатываем
шаги), если текущее решение не подходит. Это помогает эффективно обходить все возможные комбинации или перестановки,
исключая невалидные пути на каждом шаге.

### 1. Permutations

Задача: Найти все возможные перестановки элементов в массиве.

Алгоритм: Мы используем рекурсивный подход, чтобы на каждом шаге выбирать элемент, добавлять его в текущую перестановку,
и рекурсивно генерировать перестановки для оставшихся элементов.

Пример кода на Go:

func permute(nums []int) [][]int {
var result [][]int
var current []int
var backtrack func([]int)

    backtrack = func(nums []int) {
        if len(nums) == 0 {
            result = append(result, append([]int(nil), current...))
            return
        }

        for i := 0; i < len(nums); i++ {
            current = append(current, nums[i])
            numsCopy := append([]int(nil), nums[:i]...)
            numsCopy = append(numsCopy, nums[i+1:]...)
            backtrack(numsCopy)
            current = current[:len(current)-1]
        }
    }

    backtrack(nums)
    return result

}

Объяснение:

- В этом решении мы рекурсивно строим все перестановки массива.
- В каждый момент времени выбираем элемент, добавляем его в текущую перестановку, и рекурсивно обрабатываем оставшиеся
  элементы.
- После обработки каждого элемента, мы откатываем выбор, возвращая текущую перестановку в исходное состояние.

### 2. Combination Sum

Задача: Найти все возможные комбинации чисел из массива, которые суммируются до целевого значения.

Алгоритм: Рекурсивно строим все возможные комбинации чисел, начиная с текущего индекса, чтобы избежать повторений.

Пример кода на Go:

func combinationSum(candidates []int, target int) [][]int {
var result [][]int
var current []int

    var backtrack func(int, int)
    backtrack = func(start, target int) {
        if target == 0 {
            result = append(result, append([]int(nil), current...))
            return
        }

        for i := start; i < len(candidates); i++ {
            if candidates[i] > target {
                continue
            }

            current = append(current, candidates[i])
            backtrack(i, target - candidates[i])  // Позволяем использовать тот же элемент несколько раз
            current = current[:len(current)-1]
        }
    }

    backtrack(0, target)
    return result

}

Объяснение:

- Мы рекурсивно перебираем все возможные комбинации чисел, начиная с текущего индекса (чтобы позволить использовать
  одинаковые элементы несколько раз).
- Когда сумма текущей комбинации равна целевому числу, добавляем эту комбинацию в результат.
- На каждом шаге откатываем последний добавленный элемент.

### 3. Subsets

Задача: Найти все возможные подмножества массива.

Алгоритм: Мы можем решать эту задачу с помощью поиска с возвратом, генерируя подмножества для каждого элемента, начиная
с текущего состояния.

Пример кода на Go:

func subsets(nums []int) [][]int {
var result [][]int
var current []int

    var backtrack func(int)
    backtrack = func(start int) {
        result = append(result, append([]int(nil), current...))

        for i := start; i < len(nums); i++ {
            current = append(current, nums[i])
            backtrack(i + 1) // не позволяю использовать тот же элемент дважды
            current = current[:len(current)-1]
        }
    }

    backtrack(0)
    return result

}

Объяснение:

- Здесь мы генерируем все возможные подмножества, начиная с пустого набора, и для каждого элемента решаем, добавить его
  в текущее подмножество или нет.
- Мы используем backtrack(i + 1), чтобы элементы не повторялись в подмножестве.

### 4. Generate Parentheses

Задача: Генерировать все возможные комбинации правильных скобок.

Алгоритм: Для каждой позиции мы можем добавить либо открывающую, либо закрывающую скобку, при этом следим за тем, чтобы
количество закрывающих скобок никогда не превышало количества открывающих.

Пример кода на Go:

func generateParenthesis(n int) []string {
var result []string
var current string

    var backtrack func(open, close int)
    backtrack = func(open, close int) {
        if len(current) == 2*n {
            result = append(result, current)
            return
        }

        if open < n {
            current += "("
            backtrack(open + 1, close)
            current = current[:len(current)-1]
        }

        if close < open {
            current += ")"
            backtrack(open, close + 1)
            current = current[:len(current)-1]
        }
    }

    backtrack(0, 0)
    return result

}

Объяснение:

- Рекурсивно строим строку, добавляя открывающую скобку, если их количество меньше n, и добавляя закрывающую скобку,
  если количество закрывающих меньше количества открывающих.
- Когда длина строки равна 2*n, мы добавляем текущую строку в результат.

### 5. Remove Invalid Parentheses

Задача: Удалить минимальное количество скобок, чтобы строка стала валидной.

Алгоритм: Используем поиск с возвратом, чтобы попробовать все возможные удаления скобок и проверить, является ли строка
валидной.

Пример кода на Go:

func removeInvalidParentheses(s string) []string {
var result []string
var current string
var visited = make(map[string]bool)

    var backtrack func(start, left, right int)
    backtrack = func(start, left, right int) {
        if start == len(s) {
            if left == right {
                if !visited[current] {
                    visited[current] = true
                    result = append(result, current)
                }
            }
            return
        }

        if s[start] == '(' {
            backtrack(start + 1, left + 1, right)
        } else if s[start] == ')' {
            backtrack(start + 1, left, right + 1)
        }

        current += string(s[start])
        backtrack(start + 1, left, right)
        current = current[:len(current)-1]
    }

    backtrack(0, 0, 0)
    return result

}

Объяснение:

- Мы пытаемся удалить скобки по одному символу, рекурсивно строя все возможные варианты строки.
- Проверяем, что полученная строка является валидной, если количество открывающих и закрывающих скобок одинаково.

### 6. N-Queens

Задача: Разместить n ферзей на шахматной доске размером n x n, чтобы они не угрожали друг другу.

Алгоритм: Используем рекурсию для размещения ферзей на доске и проверяем, не угрожают ли они друг другу по вертикали,
диагоналям и строкам.

Пример кода на Go:

func solveNQueens(n int) [][]string {
result := [][]string{}
board := make([][]string, n)

    for i := 0; i < n; i++ {
        board[i] = make([]string, n)
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }

    var backtrack func(int)
    backtrack = func(row int) {
        if row == n {
            var solution []string
            for _, r := range board {
                solution = append(solution, strings.Join(r, ""))
            }
            result = append(result, solution)
            return
        }

        for col := 0; col < n; col++ {
            if isValid(board, row, col) {
                board[row][col] = "Q"
                backtrack(row + 1)
                board[row][col] = "."
            }
        }
    }

    backtrack(0)
    return result

}

func isValid(board [][]string, row, col int) bool {
for i := 0; i < row; i++ {
if board[i][col] == "Q" {
return false
}
if col-(row-i) >= 0 && board[i][col-(row-i)] == "Q" {
return false
}
if col+(row-i) < len(board) && board[i][col+(row-i)] == "Q" {
return false
}
}
return true
}

Объяснение:

- Мы рекурсивно пробуем разместить ферзя на каждой строке, проверяя на каждом шаге, что ферзь не угрожает другим.
- После размещения всех ферзей мы добавляем решение в список.

### 7. N-Queens II

Задача: Найти количество решений для задачи N-Queens.

Алгоритм: Похож на предыдущий, но вместо сохранения всех решений мы просто подсчитываем количество валидных размещений
ферзей.

Пример кода на Go:

```go
func isValid(board [][]string, row, col int) bool {
// Проверка для текущего столбца и диагоналей
for i := 0; i < row; i++ {
if board[i][col] == "Q" { // Проверка по столбцу
return false
}
if col-(row-i) >= 0 && board[i][col-(row-i)] == "Q" { // Проверка по первой диагонали
return false
}
if col+(row-i) < len(board) && board[i][col+(row-i)] == "Q" { // Проверка по второй диагонали
return false
}
}
return true
}

func totalNQueens(n int) int {
result := 0
board := make([][]string, n)

for i := 0; i < n; i++ {
board[i] = make([]string, n)
for j := 0; j < n; j++ {
board[i][j] = "."
}
}

var backtrack func (int)
backtrack = func (row int) {
if row == n {
result++
return
}

for col := 0; col < n; col++ {
if isValid(board, row, col) {
board[row][col] = "Q"
backtrack(row + 1)
board[row][col] = "."
}
}
}

backtrack(0)
return result
}
````

Объяснение:

- Мы продолжаем использовать тот же принцип, что и в предыдущем примере, но в этот раз просто увеличиваем счетчик при
  нахождении решения.

Заключение

Поиск с возвратом — это мощный инструмент для решения задач, где мы исследуем все возможные варианты и откатываемся,
когда текущий вариант не является решением. Это позволяет решать такие задачи, как перестановки, комбинации и размещения
ферзей, эффективно и без лишних вычислений.

## Дополнительно

Давайте подробно рассмотрим три ключевых алгоритма и структуры данных: DFS / BFS, префиксное дерево (Trie) и суффиксное
дерево, а также приведем примеры на Go.

### 1. DFS / BFS (Поиск в глубину и по ширине)

Поиск в глубину (DFS)

Поиск в глубину (Depth First Search, DFS) — это алгоритм обхода графа, при котором исследуются все пути из текущей
вершины, прежде чем вернуться назад.

Пример DFS на Go:

```go
package main

import "fmt"

// DFS с использованием рекурсии
func dfs(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	fmt.Println(node)

	for _, neighbor := range graph[node] {
		dfs(graph, neighbor, visited)
	}
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	visited := make(map[int]bool)
	dfs(graph, 1, visited) // Output: 1 2 4 5 3 6
}
````

Объяснение:

- Мы начинаем с вершины 1 и рекурсивно исследуем все ее соседей (первый уровень), затем переходим к соседям этих
  соседей (второй уровень), пока не посетим все вершины графа.
- Структура данных visited отслеживает посещенные вершины, чтобы избежать зацикливания.

Поиск в ширину (BFS)

Поиск в ширину (Breadth First Search, BFS) — это алгоритм обхода графа, при котором сначала исследуются все соседние
вершины текущей вершины, а затем их соседи и так далее.

Пример BFS на Go:

```go
package main

import "fmt"

func bfs(graph map[int][]int, start int) {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	bfs(graph, 1) // Output: 1 2 3 4 5 6
}
````

Объяснение:

- В BFS мы используем очередь (queue) для обработки вершин. Сначала обрабатываем вершину, затем добавляем все ее соседей
  в очередь для дальнейшей обработки.
- Алгоритм обрабатывает граф по уровням (сначала все вершины на уровне 1, затем все на уровне 2 и так далее).

### 2. Префиксное дерево (Trie)

Префиксное дерево (Trie) — это структура данных для хранения строк, где каждый узел представляет собой общий префикс для
нескольких строк. Это позволяет эффективно искать слова с общими префиксами.

Пример на Go:

```go
package main

import "fmt"

// Узел дерева
type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

// Структура самого префиксного дерева
type Trie struct {
	root *TrieNode
}

// Функция для создания нового Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Вставка слова в Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEndOfWord = true
}

// Поиск слова в Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEndOfWord
}

// Поиск по префиксу
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: false
	fmt.Println(trie.StartsWith("app")) // Output: true
}
````

Объяснение:

- Мы создаем структуру данных Trie, где каждый узел может содержать несколько дочерних узлов (по одному на каждый символ
  алфавита).
- Функции Insert и Search позволяют вставлять строки и проверять их наличие в дереве.
- StartsWith проверяет, начинается ли какое-либо слово в дереве с указанного префикса.

### 3. Суффиксное дерево

Суффиксное дерево — это структура данных, которая представляет все суффиксы строки. Суффиксное дерево полезно для
быстрого поиска подстрок в строках, поиска повторяющихся подстрок, решения задачи на наибольший общий префикс и других
задач, связанных с текстом.

Реализация суффиксного дерева достаточно сложная, и она включает в себя построение дерева, где каждый путь от корня к
листу представляет собой суффикс строки.

Пример суффиксного дерева в Go выходит за пределы базовой реализации, так как она включает сложные операции по
построению и манипулированию деревом. Однако, общее описание и подход следующий:

1. Построение дерева: Для каждой позиции строки мы создаем путь, который представляет суффикс строки, начиная с этой
   позиции.
2. Поиск в дереве: Для поиска подстроки мы ищем этот суффикс в дереве.
3. Оптимизация: Для уменьшения памяти и улучшения времени выполнения можно использовать методы, такие как сжижение
   суффиксных ссылок (suffix link compression).

Пример базовой реализации суффиксного дерева в Go может выглядеть следующим образом:

```go
package main

import "fmt"

type SuffixTreeNode struct {
	children   map[rune]*SuffixTreeNode
	suffixLink *SuffixTreeNode
	start      int
	end        *int
}

type SuffixTree struct {
	root *SuffixTreeNode
	text string
}

func NewSuffixTree(text string) *SuffixTree {
	tree := &SuffixTree{root: &SuffixTreeNode{children: make(map[rune]*SuffixTreeNode)}}
	tree.text = text
	// Алгоритм для построения дерева здесь
	return tree
}

func (tree *SuffixTree) Build() {
	// Алгоритм для построения суффиксного дерева
	// Например, алгоритм Укконе
}

func main() {
	text := "banana"
	tree := NewSuffixTree(text)
	tree.Build()
	fmt.Println("Suffix tree built for:", text)
}
````

Объяснение:

- Суффиксное дерево строится на основе алгоритмов, таких как алгоритм Укконе, который строит дерево за линейное время.
- В примере выше создается узел дерева с полями children (для хранения детей), suffixLink (ссылка на суффиксный
  суффикс), и start и end для определения диапазона символов в тексте.

### 4. Quad Tree (Квадратное дерево)

Алгоритм Quad Tree — это структура данных, используемая для представления двухмерных пространств. Каждый узел дерева
может иметь до четырех дочерних узлов, что позволяет эффективно хранить данные и выполнять операции в двумерных
пространствах, таких как изображения или географические данные.

Применение Quad Tree:

- Используется для эффективной работы с большими двухмерными данными, например, для хранения и поиска в изображениях,
  картографических данных или в обработке геометрии.
- Часто применяется в играх для управления пространством и физическими симуляциями.

Пример реализации Quad Tree на Go:

```go
package main

import "fmt"

type Point struct {
	x, y int
}

type QuadTree struct {
	TopLeft     *QuadTree
	TopRight    *QuadTree
	BottomLeft  *QuadTree
	BottomRight *QuadTree
	Point       *Point
	IsLeaf      bool
}

func NewQuadTree(x, y int) *QuadTree {
	return &QuadTree{
		Point:  &Point{x, y},
		IsLeaf: true,
	}
}

// Вставка точки в QuadTree
func (tree *QuadTree) Insert(x, y int) {
	if tree.IsLeaf {
		if tree.Point == nil {
			tree.Point = &Point{x, y}
		} else {
			// Разделяем узел на четыре
			tree.split()
			tree.Insert(tree.Point.x, tree.Point.y)
			tree.Insert(x, y)
		}
	} else {
		if x < tree.Point.x {
			if y < tree.Point.y {
				tree.BottomLeft.Insert(x, y)
			} else {
				tree.TopLeft.Insert(x, y)
			}
		} else {
			if y < tree.Point.y {
				tree.BottomRight.Insert(x, y)
			} else {
				tree.TopRight.Insert(x, y)
			}
		}
	}
}

// Разделение узла на четыре
func (tree *QuadTree) split() {
	tree.IsLeaf = false
	tree.TopLeft = NewQuadTree(tree.Point.x-1, tree.Point.y-1)
	tree.TopRight = NewQuadTree(tree.Point.x+1, tree.Point.y-1)
	tree.BottomLeft = NewQuadTree(tree.Point.x-1, tree.Point.y+1)
	tree.BottomRight = NewQuadTree(tree.Point.x+1, tree.Point.y+1)
}

func main() {
	qt := NewQuadTree(0, 0)
	qt.Insert(1, 1)
	qt.Insert(2, 2)
	fmt.Println("QuadTree built successfully!")
}
````

Объяснение:

- В этом примере, если мы вставляем точку в Quad Tree, дерево будет разделяться на 4 части, когда нужно вставить более
  одной точки.
- Каждую точку можно вставить в одну из четырех областей: верхнюю левую, верхнюю правую, нижнюю левую или нижнюю правую.

### 5. Дерево отрезков (Segment Tree)

Алгоритм Segment Tree — это структура данных, предназначенная для решения задач, связанных с диапазонами (отрезками) на
числовых данных. Например, для нахождения суммы элементов на отрезке, минимального или максимального значения.

Применение Segment Tree:

- Часто используется для обработки запросов на отрезках (например, нахождение суммы на диапазоне или минимального
  значения на диапазоне).
- Позволяет эффективно обновлять данные и быстро отвечать на запросы.

Пример реализации Segment Tree на Go:

````go
package main

import "fmt"

// Структура для сегментного дерева
type SegmentTree struct {
	tree []int
	n    int
}

// Конструктор сегментного дерева
func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	tree := make([]int, 4*n)
	st := &SegmentTree{tree: tree, n: n}
	st.build(arr, 0, 0, n-1)
	return st
}

// Рекурсивное построение сегментного дерева
func (st *SegmentTree) build(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
	} else {
		mid := (start + end) / 2
		st.build(arr, 2*node+1, start, mid)
		st.build(arr, 2*node+2, mid+1, end)
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

// Запрос суммы на отрезке [L, R]
func (st *SegmentTree) query(node, start, end, L, R int) int {
	// Полностью выходит за пределы
	if R < start || end < L {
		return 0
	}
	// Полностью входит в пределы
	if L <= start && end <= R {
		return st.tree[node]
	}
	// Часть пересекает
	mid := (start + end) / 2
	left := st.query(2*node+1, start, mid, L, R)
	right := st.query(2*node+2, mid+1, end, L, R)
	return left + right
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	// Запрос суммы элементов на отрезке [1, 3]
	fmt.Println(st.query(0, 0, len(arr)-1, 1, 3)) // Вывод: 15 (3 + 5 + 7)
}
````

Объяснение:

- Дерево строится рекурсивно, разделяя массив на подмассивы. Каждый узел дерева хранит сумму элементов в соответствующем
  диапазоне.
- query выполняет запрос на отрезке, возвращая сумму элементов на этом отрезке.

### 6. Система непересекающихся множеств (Union-Find)

Алгоритм Union-Find (также известен как Disjoint Set Union, DSU) — это структура данных, предназначенная для работы с
непересекающимися множествами. Эта структура эффективно поддерживает операции объединения (Union) и нахождения (Find),
что делает ее полезной для задач, таких как нахождение компонент связности в графах или объединение данных.

Применение Union-Find:

- Используется для решения задач на графах, например, для нахождения компонент связности.
- Применяется в алгоритмах, таких как алгоритм Краскала для нахождения минимального остовного дерева.

Пример реализации Union-Find на Go:

```go
package main

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// Union by rank
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func main() {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	fmt.Println(uf.Find(0)) // Output: 0
	fmt.Println(uf.Find(2)) // Output: 0 (they are connected)
	uf.Union(3, 4)
	fmt.Println(uf.Find(3)) // Output: 3
}
````

Объяснение:

- Find — находит корень множества, к которому принадлежит элемент. Мы используем сжатие путей, чтобы улучшить
  эффективность.
- Union — объединяет два множества. Используем объединение по рангу, чтобы минимизировать глубину дерева и улучшить
  производительность.

### 7. Алгоритмы Ли, Дейкстры и Флойда-Уоршела

Алгоритм Ли (Lee Algorithm) — это алгоритм для нахождения кратчайшего пути в невзвешенных графах (или с одинаковыми
весами рёбер). Обычно используется в задачах, связанных с лабиринтами или сетями, где граф не имеет весов или все рёбра
имеют одинаковую стоимость.

Алгоритм Дейкстры (Dijkstra’s Algorithm) — это алгоритм для нахождения кратчайшего пути в графе с положительными весами
рёбер. Алгоритм работает с графами, где все веса рёбер положительные.

Алгоритм Флойда-Уоршела (Floyd-Warshall Algorithm) — это алгоритм для нахождения всех кратчайших путей между всеми
парами вершин в графе. Этот алгоритм работает и с графами, содержащими отрицательные веса рёбер, при условии, что граф
не содержит отрицательных циклов.

#### 7.1 Алгоритм Дейкстры (Dijkstra’s Algorithm)

Алгоритм Дейкстры использует жадный подход для нахождения кратчайшего пути от начальной вершины ко всем остальным
вершинам в графе с положительными весами рёбер. Он гарантирует, что в процессе работы будут найдены минимальные пути для
каждой вершины.

Пример реализации алгоритма Дейкстры на Go:

```go
package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node, weight int
}

type Graph struct {
	adjList [][]Edge
}

func NewGraph(n int) *Graph {
	adjList := make([][]Edge, n)
	return &Graph{adjList: adjList}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.adjList[u] = append(g.adjList[u], Edge{v, w})
	g.adjList[v] = append(g.adjList[v], Edge{u, w})
}

type MinHeap struct {
	nodes []*Node
}

type Node struct {
	vertex, dist int
}

func (h *MinHeap) Len() int { return len(h.nodes) }

func (h *MinHeap) Less(i, j int) bool { return h.nodes[i].dist < h.nodes[j].dist }

func (h *MinHeap) Swap(i, j int) { h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i] }

func (h *MinHeap) Push(x interface{}) {
	h.nodes = append(h.nodes, x.(*Node))
}

func (h *MinHeap) Pop() interface{} {
	old := h.nodes
	n := len(old)
	x := old[n-1]
	h.nodes = old[0 : n-1]
	return x
}

func Dijkstra(g *Graph, start int) []int {
	dist := make([]int, len(g.adjList))
	for i := range dist {
		dist[i] = int(1e9) // Используем большое значение для "бесконечности"
	}
	dist[start] = 0

	pq := &MinHeap{}
	heap.Push(pq, &Node{start, 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Node)

		for _, edge := range g.adjList[u.vertex] {
			v := edge.node
			weight := edge.weight
			if dist[u.vertex]+weight < dist[v] {
				dist[v] = dist[u.vertex] + weight
				heap.Push(pq, &Node{v, dist[v]})
			}
		}
	}

	return dist
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 2, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(2, 3, 3)
	g.AddEdge(1, 3, 5)

	dist := Dijkstra(g, 0)
	fmt.Println("Distances:", dist)
}
````

Объяснение:

- В этом примере граф представлен с помощью списка смежности, а для поиска кратчайших путей используется минимальная
  куча для эффективного извлечения вершины с наименьшим расстоянием.
- Алгоритм Дейкстры гарантирует нахождение кратчайшего пути от начальной вершины до всех остальных в графе с
  положительными весами рёбер.

#### 7.2 Алгоритм Флойда-Уоршела (Floyd-Warshall Algorithm)

Алгоритм Флойда-Уоршела находит кратчайшие пути между всеми парами вершин в графе, что делает его полезным для решения
задач, связанных с анализом всех возможных путей в графе.

Пример реализации алгоритма Флойда-Уоршела на Go:

```go
package main

import "fmt"

const INF = 1000000000

func FloydWarshall(graph [][]int) [][]int {
	n := len(graph)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func main() {
	graph := [][]int{
		{0, 3, INF, INF},
		{2, 0, INF, 1},
		{INF, 7, 0, 2},
		{INF, INF, 1, 0},
	}

	dist := FloydWarshall(graph)
	fmt.Println("Shortest distances between all pairs of vertices:")
	for _, row := range dist {
		fmt.Println(row)
	}
}
````

Объяснение:

- В этом примере используется матрица смежности для представления графа, а алгоритм Флойда-Уоршела находит кратчайшие
  пути между всеми парами вершин, обновляя матрицу расстояний.

### 8. Топологическая сортировка

Алгоритм топологической сортировки используется для сортировки направленных ацикличных графов (DAG). Он упорядочивает
вершины таким образом, что для каждой пары смежных вершин, вершина-источник идет перед вершиной-целью.

Пример топологической сортировки на Go:

```go
package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func TopologicalSort(g *Graph) []int {
	visited := make(map[int]bool)
	stack := []int{}
	var dfs func(int)

	dfs = func(v int) {
		visited[v] = true
		for _, neighbor := range g.adjList[v] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		stack = append([]int{v}, stack...)
	}

	for node := range g.adjList {
		if !visited[node] {
			dfs(node)
		}
	}

	return stack
}

func main() {
	g := NewGraph()
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	sorted := TopologicalSort(g)
	fmt.Println("Topological Sort:", sorted)
}
````

Объяснение:

- В этом примере используется рекурсивный DFS для обхода графа. После завершения обработки вершины, она добавляется в
  стек, что позволяет получить топологически отсортированные вершины.

### 9. Алгоритм Кадане (Kadane’s Algorithm)

Алгоритм Кадане находит максимальную сумму подмассива в массиве с помощью динамического программирования за время O(n).

Пример алгоритма Кадане на Go:

```go
package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum subarray sum:", maxSubArray(nums))
}
````

Объяснение:

- Алгоритм Кадане проходит по массиву и для каждого элемента обновляет текущую сумму подмассива. Если сумма становится
  отрицательной, начинаем новый подмассив. Максимальная сумма хранится в переменной maxSum.

### 10. Алгоритм Кнута-Мориса-Пратта (KMP Algorithm)

Алгоритм Кнута-Мориса-Пратта (KMP) используется для поиска подстроки в строке за время O(n + m), где n — длина основной
строки, а m — длина подстроки.

Пример алгоритма КМП на Go:

```go
package main

import "fmt"

func KMP(text, pattern string) int {
	lps := computeLPS(pattern)
	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == len(pattern) {
			return i - j
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func main() {
	text := "abxabcabcaby"
	pattern := "abcaby"
	fmt.Println("Pattern found at index:", KMP(text, pattern))
}
```

Объяснение:

- Алгоритм КМП использует вспомогательный массив LPS (longest prefix suffix), чтобы избегать ненужных сравнений при
  поиске подстроки, делая поиск быстрее.

Эти алгоритмы широко используются в различных областях, таких как поиск путей в графах, обработка строк, и анализ
данных.

### 11. Динамическое программирование

Динамическое программирование (DP) — это метод решения задач, который использует разбиение задачи на подзадачи и
запоминает результаты уже решённых подзадач, чтобы избежать повторных вычислений.

Основные принципы DP:

1. Определение подзадач — разбиваем исходную задачу на меньшие подзадачи.
2. Запоминание промежуточных результатов — сохраняем решения подзадач в таблице для использования в будущем.
3. Оптимизация — стараемся избегать перерасчёта тех же подзадач.

Пример: Задача о рюкзаке (0/1 Knapsack)

Дано множество предметов, каждый с определённым весом и стоимостью, и рюкзак с ограничением по весу. Нужно выбрать,
какие предметы взять в рюкзак, чтобы максимизировать стоимость, не превышая ограничение по весу.

```go
package main

import "fmt"

func knapsack(weights []int, values []int, W int) int {
	n := len(weights)
	dp := make([][]int, n+1)

	// Инициализируем таблицу DP
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	// Заполнение таблицы
	for i := 1; i <= n; i++ {
		for w := 1; w <= W; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	W := 5 // Максимальный вес рюкзака

	fmt.Println("Max value:", knapsack(weights, values, W))
}
```

Объяснение:

- В этом примере мы используем динамическое программирование для решения задачи о рюкзаке. Создаём двумерную таблицу dp,
  где dp[i][w] — это максимальная стоимость, которую можно получить, используя первые i предметов и не превышая вес w.
  Мы поочередно заполняем таблицу, основываясь на предыдущих результатах.

### 12. Quick Select

Quick Select — это алгоритм для поиска k-ого по величине элемента в неотсортированном массиве. Это модификация алгоритма
быстрой сортировки (Quick Sort), но вместо сортировки всего массива, мы выполняем только часть работы, чтобы найти
нужный элемент.

Пример Quick Select:

```go
package main

import "fmt"

func quickSelect(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	pivotIndex := partition(arr, left, right)
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect(arr, left, pivotIndex-1, k)
	} else {
		return quickSelect(arr, pivotIndex+1, right, k)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := quickSelect(arr, 0, len(arr)-1, k-1) // k-1 для 0-индексации
	fmt.Println("The", k, "th largest element is:", result)
}
```

Объяснение:

- В этом примере алгоритм Quick Select используется для нахождения k-го по величине элемента. Мы используем метод
  partition из быстрой сортировки для выбора опорного элемента и рекурсивно ищем нужный элемент в одном из подмассивов.

### 13. Дерево Меркла (Merkle Tree)

Дерево Меркла — это бинарное дерево, в котором каждый листовой узел содержит хэш данных, а каждый промежуточный узел
содержит хэш от объединённых хэшей его дочерних узлов. Дерево Меркла часто используется в криптовалютных системах и
распределённых системах для проверки целостности данных.

Пример дерева Меркла:

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Hash  []byte
	Left  *MerkleNode
	Right *MerkleNode
}

func NewMerkleTree(data []string) *MerkleTree {
	var nodes []MerkleNode
	for _, d := range data {
		nodes = append(nodes, MerkleNode{Hash: sha256.Sum256([]byte(d))[:32]})
	}

	for len(nodes) > 1 {
		var newLevel []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			var right MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}
			left := nodes[i]
			combined := append(left.Hash, right.Hash...)
			hash := sha256.Sum256(combined)
			newLevel = append(newLevel, MerkleNode{Hash: hash[:32], Left: &left, Right: &right})
		}
		nodes = newLevel
	}

	return &MerkleTree{Root: &nodes[0]}
}

func main() {
	data := []string{"a", "b", "c", "d"}
	tree := NewMerkleTree(data)

	fmt.Printf("Merkle Root Hash: %x\n", tree.Root.Hash)
}
```

Объяснение:

- В примере создаётся дерево Меркла, где каждый узел хранит хэш данных. При объединении двух дочерних узлов их хэши
  соединяются, и получается хэш родительского узла. В итоге, корневой узел дерева Меркла является хэшом всех данных.

### 14. Персистентные структуры данных (Persistent Data Structures)

Персистентные структуры данных — это структуры данных, которые сохраняют свои прошлые версии, когда они изменяются. В
отличие от обычных структур данных, где изменения модифицируют данные напрямую, персистентные структуры сохраняют
неизменность старых версий.

Пример реализации персистентного стека:

```go
package main

import "fmt"

// Стек, который сохраняет все версии
type PersistentStack struct {
	value int
	prev  *PersistentStack
}

func NewPersistentStack() *PersistentStack {
	return nil
}

func (s *PersistentStack) Push(value int) *PersistentStack {
	return &PersistentStack{value, s}
}

func (s *PersistentStack) Pop() (*PersistentStack, int) {
	if s == nil {
		return nil, 0
	}
	return s.prev, s.value
}

func main() {
	stack := NewPersistentStack()
	stack1 := stack.Push(1)
	stack2 := stack1.Push(2)
	stack3 := stack2.Push(3)

	fmt.Println("Top of stack3:", stack3.value)

	stack2, value := stack3.Pop()
	fmt.Println("Popped value:", value)
	fmt.Println("Top of stack2:", stack2.value)
}
```

Объяснение:

- В этом примере реализован персистентный стек, где каждый новый элемент создаёт новую версию стека, оставляя старые
  версии неизменными. Когда вызывается метод Pop, он возвращает текущую версию стека и её верхний элемент.

### 15. Фильтр Блума (Bloom Filter)

Фильтр Блума — это вероятностная структура данных, предназначенная для тестирования, принадлежит ли элемент множеству.
Он может возвращать два результата:

1. True — элемент, возможно, присутствует в множестве.
2. False — элемент точно отсутствует в множестве.

Фильтр Блума использует несколько хеш-функций и битовый массив, чтобы эффективно проверять наличие элементов, но с
возможностью ошибок (ложноположительных срабатываний).

Пример фильтра Блума на Go:

```go
package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray  []bool
	size      int
	hashFuncs []func(string) int
}

func NewBloomFilter(size int, hashFuncs []func(string) int) *BloomFilter {
	return &BloomFilter{
		bitArray:  make([]bool, size),
		size:      size,
		hashFuncs: hashFuncs,
	}
}

func (bf *BloomFilter) Add(item string) {
	for _, hashFunc := range bf.hashFuncs {
		index := hashFunc(item) % bf.size
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	for _, hashFunc := range bf.hashFuncs {
		index := hashFunc(item) % bf.size
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

func simpleHash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	hashFuncs := []func(string) int{simpleHash}
	bf := NewBloomFilter(100, hashFuncs)

	bf.Add("apple")
	bf.Add("banana")

	fmt.Println(bf.Contains("apple"))  // true
	fmt.Println(bf.Contains("banana")) // true
	fmt.Println(bf.Contains("cherry")) // false
}
```

Объяснение:

- Мы создаём фильтр Блума с помощью битового массива и множества хеш-функций.
- Для каждого добавленного элемента вызываются хеш-функции, которые изменяют биты в массиве.
- При проверке, если хотя бы один бит для элемента равен 0, то элемент точно отсутствует. Если все биты равны 1, элемент
  может присутствовать (с вероятностью ложноположительного результата).

### 16. Битовая карта (Bit Map)

Битовая карта (или битовое поле) — это структура данных, которая использует массив битов (0 или 1) для представления
большого множества данных. Каждый бит в карте может представлять наличие или отсутствие элемента, что позволяет
эффективно использовать память.

Пример битовой карты на Go:

```go
package main

import "fmt"

type BitMap struct {
	bits []uint64
}

func NewBitMap(size int) *BitMap {
	return &BitMap{
		bits: make([]uint64, (size+63)/64), // 64 бита в одном uint64
	}
}

func (bm *BitMap) Set(index int) {
	bm.bits[index/64] |= 1 << (index % 64)
}

func (bm *BitMap) Get(index int) bool {
	return bm.bits[index/64]&(1<<(index%64)) != 0
}

func main() {
	bm := NewBitMap(100)
	bm.Set(10)
	bm.Set(30)

	fmt.Println("Is 10 set?", bm.Get(10)) // true
	fmt.Println("Is 20 set?", bm.Get(20)) // false
	fmt.Println("Is 30 set?", bm.Get(30)) // true
}
```

Объяснение:

- В этом примере мы используем массив bits типа uint64, чтобы представить множество из 100 элементов.
- Функция Set устанавливает бит на заданном индексе, а функция Get проверяет, установлен ли бит.
- Каждое значение в битовой карте занимает 64 бита (1 uint64), что позволяет эффективно хранить большое количество
  данных.

### 17. Дерево B / B+ (B-tree / B+ tree)

Дерево B — это самобалансирующееся дерево поиска, которое хранит отсортированные данные и позволяет эффективно выполнять
операции вставки, удаления и поиска. Это дерево используется в базах данных и файловых системах для быстрого поиска на
дисках.

B+ дерево — это разновидность дерева B, в котором все данные находятся только в листовых узлах, а внутренние узлы
содержат только ключи.

Пример B-дерева:

Для более полного примера реализации B или B+ дерева нужно будет подробно проделывать операции вставки, удаления и
балансировки, что выходит за пределы простого объяснения. Вместо этого приведём концептуальное описание:

- B-дерево — это дерево, где каждый узел может иметь более двух дочерних элементов (в отличие от бинарного дерева).
- B+ дерево — расширение B-дерева, где только листовые узлы содержат данные, а внутренние узлы используются только для
  навигации.

Операции вставки и поиска выполняются за время O(log n), где n — количество элементов в дереве.

### 18. LSM-дерево (Log-Structured Merge Tree)

LSM-дерево — это структура данных, использующаяся для записи данных в базы данных с высоким уровнем записи. Оно
эффективно работает с данными, которые часто изменяются, например, в случае логов или других записей.

LSM-дерево работает, разделяя данные на два уровня:

1. MemTable — это структура данных в памяти, куда сначала записываются новые данные.
2. SSTable (Sorted String Table) — на диск сохраняются данные в отсортированном виде.

При достижении определённого порога данных в MemTable они сбрасываются на диск в виде SSTable, и позже происходит их
слияние и упорядочивание.

Пример простого LSM-дерева:

Пример LSM-дерева будет включать операции записи в MemTable, сброс на диск в SSTable и слияние. Простейшая версия
LSM-дерева выходит за пределы данного ответа, но принцип работы следующий:
- Вначале данные записываются в память (MemTable).
- Когда MemTable достигает определённого размера, она сбрасывается в SSTable.
- Процесс слияния SSTable происходит периодически для оптимизации.

LSM-дерево используется в таких системах, как LevelDB и RocksDB, для эффективной обработки больших объёмов данных.

### 19. Splay-дерево (Splay Tree)
Splay-дерево — это самобалансирующееся бинарное дерево поиска, которое использует операции “поднятия” (splaying) для
того, чтобы часто используемые элементы оказывались ближе к корню дерева. Это достигается за счёт выполнения операций
вращений (zig, zig-zig, и т.д.) после каждого доступа к элементу.

Особенности Splay-дерева:
- После каждой операции поиска, вставки или удаления дерево восстанавливает свою структуру таким образом, что последний
  доступный элемент оказывается в корне дерева. Это помогает ускорить дальнейшие операции с этим элементом.
- Нет необходимости хранить балансировку дерева в явном виде (как в AVL или красно-чёрных деревьях), что упрощает его
  реализацию.

Пример Splay-дерева на Go:

```go
package main

import "fmt"

// Определение структуры узла дерева
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Функция для выполнения вращения влево
func rotateLeft(root *Node) *Node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	return newRoot
}

// Функция для выполнения вращения вправо
func rotateRight(root *Node) *Node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	return newRoot
}

// Функция для выполнения "splaying" на основе ключа
func splay(root *Node, key int) *Node {
	if root == nil || root.key == key {
		return root
	}

	if key < root.key {
		if root.left == nil {
			return root
		}
		if key < root.left.key {
			// Zig-zig (вправо-вправо)
			root = rotateRight(root)
			root = rotateRight(root)
		} else if key > root.left.key {
			// Zig-zag (вправо-влево)
			root.left = rotateLeft(root.left)
			root = rotateRight(root)
		}
	} else {
		if root.right == nil {
			return root
		}
		if key > root.right.key {
			// Zig-zig (влево-влево)
			root = rotateLeft(root)
			root = rotateLeft(root)
		} else if key < root.right.key {
			// Zig-zag (влево-вправо)
			root.right = rotateRight(root.right)
			root = rotateLeft(root)
		}
	}

	return root
}

func main() {
	root := &Node{key: 10}
	root.left = &Node{key: 5}
	root.right = &Node{key: 20}
	root.left.left = &Node{key: 1}
	root.right.left = &Node{key: 15}

	// "Splay" дерева с ключом 15
	root = splay(root, 15)
	fmt.Println("Root after splay:", root.key)
}
```

Объяснение:
- Splay-дерево оптимизирует доступ к часто используемым элементам, “перемещая” их в корень дерева.
- В примере дерево сначала создается с несколькими узлами. После того как мы вызываем splay(root, 15), элемент с ключом
  15 будет перемещён в корень дерева, при этом дерево будет сбалансировано с помощью вращений.

### 20. Кольцевой буфер (Circular Buffer)
Кольцевой буфер — это структура данных, которая используется для хранения фиксированного объема данных и эффективно
работает с операциями добавления и удаления элементов в цикличном режиме. Это полезно, когда нужно организовать
обработку данных, которые поступают и обрабатываются по кругу, например, в системах с ограниченной памятью.

Пример кольцевого буфера на Go:

```go
package main

import "fmt"

type CircularBuffer struct {
	buffer []int
	head   int
	tail   int
	size   int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer: make([]int, size),
		size:   size,
	}
}

func (cb *CircularBuffer) Write(value int) {
	cb.buffer[cb.tail] = value
	cb.tail = (cb.tail + 1) % cb.size
	if cb.tail == cb.head {
		cb.head = (cb.head + 1) % cb.size // перезапись данных
	}
}

func (cb *CircularBuffer) Read() (int, bool) {
	if cb.head == cb.tail {
		return 0, false // пустой буфер
	}
	value := cb.buffer[cb.head]
	cb.head = (cb.head + 1) % cb.size
	return value, true
}

func main() {
	cb := NewCircularBuffer(5)
	cb.Write(1)
	cb.Write(2)
	cb.Write(3)

	// Чтение элементов
	for i := 0; i < 3; i++ {
		if val, ok := cb.Read(); ok {
			fmt.Println(val)
		}
	}
}
```

Объяснение:
- Кольцевой буфер используется для того, чтобы не выходить за пределы фиксированного размера массива.
- Когда данные записываются в буфер, указатели на начало и конец смещаются по кругу. Когда буфер заполняется, новые
  данные заменяют старые.

### 21. Фибоначчева куча (Fibonacci Heap)
Фибоначчева куча — это структура данных, которая поддерживает эффективные операции с минимальным элементом. Она
используется для оптимизации алгоритмов, например, алгоритма Дейкстры. Фибоначчева куча работает быстрее, чем
традиционная бинарная куча для операций слияния и уменьшения ключа.

Пример использования Фибоначчевой кучи:

Полный пример реализации Фибоначчевой кучи выходит за рамки этого ответа, однако основные операции включают:

1. Insert — вставка элемента.
2. Min — извлечение минимального элемента.
3. Union — слияние двух куч.

Операции слияния и уменьшения ключей выполняются с амортизированной сложностью O(1), а извлечение минимального
элемента — с O(log n).

### 22. Жадные алгоритмы (Greedy Algorithms)
Жадные алгоритмы — это подход к решению задач, при котором на каждом шаге выбирается наилучшее решение, исходя из
текущих данных, без учета будущих шагов. Это не всегда дает оптимальное решение, но может быть эффективным для некоторых
типов задач.

Пример жадного алгоритма — задача о размене монет.
Пример жадного алгоритма на Go (задача о размене монет):

```go
package main

import "fmt"

func coinChange(coins []int, amount int) int {
	// Массив для хранения минимального числа монет для каждой суммы
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // Инициализация значением больше суммы
	}
	dp[0] = 0 // Для суммы 0 нужно 0 монет

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1 // Невозможно набрать сумму
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("Minimum coins:", coinChange(coins, amount))
}
```

Объяснение:
- В этом примере жадный алгоритм решает задачу размена монет, пытаясь с каждым шагом использовать монеты с наибольшим
  номиналом.
- Алгоритм динамически заполняет массив, в котором хранится минимальное количество монет для каждой суммы от 0 до
  требуемой. В конце возвращается минимальное количество монет для заданной суммы.

Заключение

1. DFS и BFS являются важными алгоритмами для обхода графов, каждый из которых использует различные подходы для
   исследования вершин.
2. Префиксное дерево (Trie) — это структура данных, оптимизированная для работы со строками, позволяющая быстро
   выполнять
   операции поиска по префиксу.
3. Суффиксное дерево представляет все суффиксы строки и является мощным инструментом для работы с текстами, например,
   для
   поиска подстрок или нахождения общих префиксов.
4. Quad Tree — структура данных для эффективного представления двумерных пространств с помощью четырёх областей.
   Применяется в географических приложениях и играх.
5. Segment Tree — структура для решения задач на отрезках числовых данных, таких как поиск суммы или минимального
   значения на диапазоне.
6. Union-Find (Система непересекающихся множеств) — структура для эффективного объединения и поиска компонентов в графах
   или других задачах, где элементы можно группировать в множества.
7. Алгоритмы Ли, Дейкстры и Флойда-Уоршела
8. Топологическая сортировка
9. Алгоритм Кадане (Kadane’s Algorithm)
10. Алгоритм Кнута-Мориса-Пратта (KMP Algorithm)
11. Динамическое программирование (DP) позволяет эффективно решать задачи, разбивая их на подзадачи и запоминая
    промежуточные результаты.
12. Quick Select — это эффективный алгоритм для поиска k-го по величине элемента в массиве.
13. Дерево Меркла используется для проверки целостности данных в распределённых системах.
14. Персистентные структуры данных сохраняют старые версии данных и позволяют работать с ними без потери информации.

15. Фильтр Блума — эффективная структура для проверки принадлежности элемента множеству с вероятностью
    ложноположительного результата.
16. Битовая карта — структура данных, которая позволяет эффективно хранить данные в виде битов.
17. Дерево B/B+ — самобалансирующееся дерево, используемое для эффективного поиска и манипуляций с большими объемами
    данных.
18. LSM-дерево — используется для эффективной записи и хранения данных в базах данных с высокой нагрузкой на записи.

19. Splay-дерево — самобалансирующееся бинарное дерево поиска, которое с помощью операций “поднятия” улучшает доступ к
    часто используемым элементам.
20. Кольцевой буфер — структура данных с фиксированным размером, которая позволяет эффективно записывать и читать данные
    в циклическом порядке.
21. Фибоначчева куча — структура данных, оптимизирующая алгоритмы, такие как Дейкстра, за счёт более быстрых операций
    слияния и уменьшения ключа.
22. Жадные алгоритмы — подходы к решению задач, при которых на каждом шаге принимается локально оптимальное решение, что
    может привести к глобально оптимальному решению для некоторых задач.
   
---

   Каждый из этих алгоритмов и структур данных применим в различных областях, таких как анализ графов, обработка строк и
   текстовых данных.

   Эти структуры данных играют важную роль в решении многих проблем, связанных с эффективным обработкой данных в графах,
   изображениях, и числовых отрезках.

   Эти алгоритмы и структуры данных широко применяются в различных областях, от криптографии до оптимизации вычислений и
   работы с большими данными.

   Каждая из этих структур данных имеет свои особенности и области применения, и понимание их работы полезно для
   разработки
   эффективных алгоритмов и систем.

   Эти алгоритмы и структуры данных находят применение в реальных задачах, таких как обработка графов, управление памятью и
   оптимизация вычислений.
   