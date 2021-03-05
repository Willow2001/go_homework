package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
const count = 50000

func main() {
	num := runtime.NumCPU()//返回本地机器的逻辑cpu个数
	runtime.GOMAXPROCS(num)//多操作线程处理
	doneCh := make(chan struct{})
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			find(x)
		}(i)
	}
	close(doneCh)
	wg.Wait()
	fmt.Println("main goroutine finish")
}
func find(x int){
	var m,judge int
	judge = 0
	sqrtNum := math.Sqrt(float64(x))//利用平方根提高搜索效率
	for i := 2; i <= int(sqrtNum); i++ {
		m = x % i
		if m==0{
			judge++
		}
	}
	if judge==0 && x!=1{
		fmt.Println(x)//打印素数
	}
}


