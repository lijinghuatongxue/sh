package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "汪汪🐶~~"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "喵喵🐱~~"
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		logrus.Info(reflect.TypeOf(animals))
		fmt.Printf("animal:%#v -> 📢：%s\n", animal, animal.Speak())
	}
}
