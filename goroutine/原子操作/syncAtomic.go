package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var totalSecond uint64

func workerSecond(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 10000000; i++ {
		atomic.AddUint64(&totalSecond, i)
		//logrus.Info(fmt.Sprintf("i: %d | value: %d", i, totalSecond))
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go workerSecond(&wg)
	go workerSecond(&wg)
	wg.Wait()
	fmt.Println(totalSecond)
}
