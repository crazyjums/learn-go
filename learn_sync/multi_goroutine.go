package learn_sync

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	PageCount  = 100
	TotalCount = 56765
)

// FetchDataByPage 模拟从数据库中获取数据
func FetchDataByPage(page int) []string {
	consume := getRandNumIn300And500() //模拟网络请求
	time.Sleep(time.Duration(consume) * time.Nanosecond)

	if ((page - 1) * PageCount) > TotalCount {
		return []string{}
	}

	return []string{strconv.Itoa(page)}
}

// getRandNumIn300And500 获取300到500之间的随机数
func getRandNumIn300And500() int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成300到500之间的随机数
	randomNumber := rand.Intn(201) + 300

	return randomNumber
}

func FetchWith1Goroutine() {
	var results = make([]string, 0)
	page := 1
	start := time.Now()
	for {
		result := FetchDataByPage(page)
		if len(result) == 0 {
			break
		}
		page++
		results = append(results, result...)
	}

	fmt.Println(len(results))
	fmt.Println("cost time:", time.Since(start))
}

func FetchWith10Goroutine() {
	var results = make([]string, 0)
	var wg sync.WaitGroup
	var mu sync.Mutex

	var pageNumChan = make(chan int, 1)
	pageNumChan <- 1

	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				page := <-pageNumChan
				page++
				pageNumChan <- page
				//fmt.Printf(" %v page: %v \n", getGoroutineID(), page)

				result := FetchDataByPage(page)
				if len(result) == 0 {
					break
				}
				mu.Lock()
				results = append(results, result...)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	close(pageNumChan)

	fmt.Println(len(results))
	fmt.Println("cost time:", time.Since(start))
}

func FetchWith10Goroutine2() {
	var results = make([]string, 0)
	var wg sync.WaitGroup

	var pageResultChan = make(chan []string, 10)
	var pageNumChan = make(chan int, 1)
	pageNumChan <- 1

	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				page := <-pageNumChan
				page++
				pageNumChan <- page
				//fmt.Printf(" %v page: %v \n", getGoroutineID(), page)

				result := FetchDataByPage(page)
				if len(result) == 0 {
					break
				}
				pageResultChan <- result
			}
		}()
	}

	go func() {
		for e := range pageResultChan {
			results = append(results, e...)
		}
	}()

	wg.Wait()
	close(pageNumChan)

	fmt.Println(len(results))
	fmt.Println("cost time:", time.Since(start))
}

func wgWithDeadLock() {
	var wg sync.WaitGroup
	var pageResultChan = make(chan int, 10)

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pageResultChan <- i
		}(i)
	}
	wg.Wait()

	for e := range pageResultChan {
		fmt.Println(e)
	}
	//当通道 pageResultChan 中没有数据时，for range 循环会一直阻塞在接收数据的语句上，无法继续执行下面的代码。
	close(pageResultChan)

	fmt.Println("cost time:", time.Since(start))
}

func wgWithoutDeadLock() {
	var wg sync.WaitGroup
	var pageResultChan = make(chan int, 10)

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pageResultChan <- i
		}(i)
	}
	wg.Wait()
	close(pageResultChan) //key code

	for e := range pageResultChan {
		fmt.Println(e)
	}

	fmt.Println("cost time:", time.Since(start))
}

func wgPage() {
	var wg sync.WaitGroup

	var pageNumChan = make(chan int, 1)
	pageNumChan <- 0

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			page := <-pageNumChan

			page++
			pageNumChan <- page
		}()
	}

	wg.Wait()
	close(pageNumChan)

	fmt.Println("page num:", <-pageNumChan)

	fmt.Println("cost time:", time.Since(start))
}

func wgPage2() {
	var wg sync.WaitGroup
	dataChan := make(chan int, 1)

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 从通道接收数据
			data := <-dataChan
			fmt.Printf(" %v data: %v \n", getGoroutineID(), data)

			// 对数据进行自增操作
			data++

			// 将数据放回通道
			dataChan <- data
		}()
	}

	// 将初始数据发送到通道
	dataChan <- 0

	// 等待所有协程完成
	wg.Wait()

	// 从通道接收最终结果
	result := <-dataChan

	fmt.Println("Final result:", result)
}

func getGoroutineID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int64
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}
