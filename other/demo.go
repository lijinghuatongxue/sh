package main

var a string
var done bool

func setup() {
	a = "hello world"
	done = true
}

// 顺序一次性内存模型
func main() {
	go setup()
	for !done {
	}
	print(a)
}
