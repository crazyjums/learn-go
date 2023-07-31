package learn_slice

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(array []int) []int {
	length := len(array)
	if length == 0 || length == 1 {
		return []int{}
	}

	count := 0
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			count++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	fmt.Println(count)
	return array
}

// BubbleSortV2 冒泡排序
func BubbleSortV2(array []int) []int {
	length := len(array)
	if length == 0 || length == 1 {
		return []int{}
	}

	count := 0
	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 0; j < length-i-1; j++ {
			count++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped { // 如果没有交换,说明已经有序,直接返回
			break
		}
	}

	fmt.Println(count)
	return array
}

// SelectSort 选择排序
func SelectSort(array []int) []int {
	length := len(array)
	if length == 0 || length == 1 {
		return []int{}
	}

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}

	return array
}

// MergeSort 归并排序
func MergeSort(array []int) []int {
	length := len(array)
	if length <= 1 {
		return array
	}

	left := MergeSort(array[:length/2])
	right := MergeSort(array[length/2:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	l := 0
	r := 0
	result := make([]int, 0)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	if l < len(left) {
		result = append(result, left[l:]...)
	}
	if r < len(right) {
		result = append(result, right[r:]...)
	}

	return result
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	partition := func(arr []int) int {
		pivot := arr[0]
		left, right := 1, len(arr)-1
		for left <= right {
			for left <= right && arr[left] < pivot {
				left++
			}
			for left <= right && arr[right] > pivot {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		arr[0], arr[right] = arr[right], arr[0]
		return right
	}
	pivotIndex := partition(arr)

	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
	return arr
}

func CountingSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	// 找出最大值和最小值
	max, min := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	// 计算元素范围
	rangeSize := max - min + 1

	// 统计元素出现的次数
	count := make([]int, rangeSize)
	for i := 0; i < n; i++ {
		count[arr[i]-min]++
	}
	fmt.Println("统计元素出现的次数", count)

	// 计算累计次数
	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}
	fmt.Println("累计次数", count)

	// 排序
	sorted := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		fmt.Println("arr[i]-min", arr[i]-min, "count[arr[i]-min]-1", count[arr[i]-min]-1)
		sorted[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
		fmt.Println("sorted[count[arr[i]-min]-1]", sorted)
	}

	// 将排序结果复制回原数组
	for i := 0; i < n; i++ {
		arr[i] = sorted[i]
	}

	return arr
}
