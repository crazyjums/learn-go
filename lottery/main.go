//package main
//
//import (
//	"fmt"
//	"math/rand"
//)
//
//func main() {
//	//中奖奖品
//	prizeArr := [6]map[string]interface{}{
//		{"id": 1, "prize": "现金500W", "v": 1}, //概率为1/200
//		{"id": 2, "prize": "iphone7", "v": 5},
//		{"id": 3, "prize": "耐克跑鞋", "v": 10},
//		{"id": 4, "prize": "魔声耳机", "v": 24},
//		{"id": 5, "prize": "蓝牙音响", "v": 60},
//		{"id": 6, "prize": "现金1元", "v": 100},
//	}
//
//	/*
//	 * 对数组进行处理
//	 */
//	item := make(map[int]int)
//	for _, v := range prizeArr {
//		item[v["id"].(int)] = v["v"].(int)
//	}
//	/*
//	   map[1:1 2:5 3:10 4:24 5:60 6:100]
//	*/
//
//	getRand := func(item map[int]int) int {
//		num := 0
//		for _, v := range item {
//			num += v
//		}
//		result := 0
//		for k, v := range item {
//			rand.Seed(int64(k)) //设置随机数种子
//			randNum := rand.Intn(num) + 1
//			if randNum <= v {
//				result = k
//				break
//			} else {
//				num -= v
//			}
//		}
//		return result
//	}
//
//	times := 1000000
//	prizes := make(map[string]int)
//	for i := 0; i < times; i++ {
//		fmt.Println(i)
//		r := getRand(item)
//		p := prizeArr[r-1]["prize"].(string)
//		if _, ok := prizes[p]; ok {
//			prizes[p]++
//		} else {
//			prizes[p] = 1
//		}
//	}
//
//	fmt.Println("概率：")
//	for p, nums := range prizes {
//		fmt.Printf("%s------>%.2f\n", p, float64(nums)/float64(times))
//	}
//	fmt.Println()
//	fmt.Println()
//	for p, v := range item {
//		fmt.Printf("%s------>%.2f\n", prizeArr[p-1]["prize"].(string), float64(v)/200)
//	}
//}
//


package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 奖品结构体
type Prize struct {
	name  string  // 奖品名称
	prob  float64 // 中奖概率
	total int     // 总数
	left  int     // 剩余数量
}

// 礼物转盘结构体
type GiftWheel struct {
	prizes []Prize // 奖品列表
	r      *rand.Rand // 随机数生成器
}

// 初始化礼物转盘
func NewGiftWheel(prizes []Prize) *GiftWheel {
	wheel := &GiftWheel{
		prizes: prizes,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return wheel
}

// 抽奖方法
func (w *GiftWheel) Draw() *Prize {
	// 计算总概率
	totalProb := 0.0
	for _, prize := range w.prizes {
		totalProb += prize.prob
	}

	// 生成随机数，判断落在哪个奖品区间内
	randNum := w.r.Float64() * totalProb
	probSum := 0.0
	for i := range w.prizes {
		prize := &w.prizes[i]
		if prize.left > 0 {
			probSum += prize.prob
			if randNum <= probSum {
				prize.left--
				return prize
			}
		}
	}

	// 如果所有奖品都已抽完，则返回 nil
	return nil
}

func main() {
	// 奖品列表，每种奖品的中奖概率和总数可以进行配置
	prizes := []Prize{
		{name: "一等奖", prob: 0.1, total: 1000, left: 1000},
		{name: "二等奖", prob: 0.2, total: 2000, left: 2000},
		{name: "三等奖", prob: 0.3, total: 3000, left: 30000},
		{name: "四等奖", prob: 0.4, total: 4000, left: 4000},
	}

	// 初始化礼物转盘
	wheel := NewGiftWheel(prizes)

	var m map[string]int
	m = make(map[string]int, 10000)
	// 抽奖，重复抽取 10 次
	for i := 1; i <= 10000; i++ {
		prize := wheel.Draw()
		if prize != nil {
			if v,ok := m[prize.name]; ok {
				v++
				m[prize.name]=v
			} else {
				m[prize.name] = 1
			}
			fmt.Printf("第 %d 次抽奖，恭喜您获得了 %s！\n", i, prize.name)
		} else {
			m["未中奖"]++
			fmt.Printf("第 %d 次抽奖，很遗憾您未中奖。\n", i)
		}
	}
	fmt.Println(m)
}
