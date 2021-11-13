package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10000000; i++ {
		total.Lock()
		total.value += i
		//logrus.Info(fmt.Sprintf("i: %d | value: %d", i, total.value))
		total.Unlock()
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}
