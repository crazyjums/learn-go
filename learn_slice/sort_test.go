package learn_slice

import (
	"fmt"
	"testing"
)

func TestBobbleSort(t *testing.T) {
	ans := BubbleSort([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8})
	fmt.Println(ans)
}

func TestBobbleSortV2(t *testing.T) {
	ans := BubbleSortV2([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8})
	fmt.Println(ans)
}

func TestSelectSort(t *testing.T) {
	ans := SelectSort([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8})
	fmt.Println(ans)
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8}))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8}))
}

func TestCountingSort(t *testing.T) {
	fmt.Println(CountingSort([]int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8}))
}

func Test1(t *testing.T) {
	data := []int{11, 9, 1, 5, 3, 0, 22, 4, 6, 7, 8}
	s1 := data[2:5]
	s2 := s1[2:6:7]
	fmt.Println(s1, len(s1), cap(s1)) // [1 5 3] 3 9
	fmt.Println(s2, len(s2), cap(s2)) // [3 0 22 4] 4 5
	app(s1)
	app(s2)
	fmt.Println(s1, len(s1), cap(s1))       // [1 5 3] 3 9
	fmt.Println(s2, len(s2), cap(s2))       // [3 111 222 4] 4 5
	fmt.Println(data, len(data), cap(data)) // [11 9 1 5 3 111 222 4 111 7 8] 11 11
}

func app(s []int) {
	s = append(s, 111)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 222)
	fmt.Println(len(s), cap(s), s)
}

func Test2(t *testing.T) {
	//s := "12312vads"
	//rs := []rune(s)
	//l, r := 0, len(s)-1
	//for l < r {
	//	rs[l], rs[r] = rs[r], rs[l]
	//	l++
	//	r--
	//}
	//fmt.Println(string(rs))
	//fmt.Println(app2(100))
	fmt.Println(foo2())
}

func app2(x int) int {
	defer func() {
		x = 1000
	}()

	x++
	return x
}

func foo() {
	x := 10
	defer fmt.Println("Value of x in defer:", x) // 不合法
	x = 20
	fmt.Println("Value of x before defer:", x)
}

func foo2() (x int) {
	defer func() {
		x = 100
	}()

	x = 10
	return x
}

func Test3(t *testing.T) {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
	ans := []rune{}
	fmt.Println(len(ans), cap(ans), ans == nil)
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	for k, v := range m {
		if k == 1 {
			delete(m, k)
		}
		fmt.Println(k, v)
	}
	fmt.Println(m)
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
		i++
	}
	*/

	for i := range s {
		s[i] += 1
	}

	ch := make(chan<- int)
	fmt.Println(ch)
}
