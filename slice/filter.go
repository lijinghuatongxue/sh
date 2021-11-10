package main

import "github.com/sirupsen/logrus"

// Filter 从byte slice 中剔除指定的字节元素
func Filter(s []byte,fn func(x byte) bool)  []byte{
	b := s[:0]
	for _,x := range s{
		if !fn(x){
			b = append(b,x)
		}
	}
	return b
}
func main()  {
	a := []byte{' ',2,'2',2,98}
	result := Filter(a,fn)
	logrus.Info(result)
}

func fn(x byte) bool {
	if x == 2 {
		return true
	}
	return false
}