package main

func Fibonacci(n int) int {
	return fibIterator(n)
}

func fibGoroutine() {
	c := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < 10; i = i + 1 {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()
	for i := range c {
		println(i)
	}
	println("Finished")
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibDp(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i < n+1; i = i + 1 {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fibIterator(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}
