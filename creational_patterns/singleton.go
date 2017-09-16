package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Instance interface {
	TestFunc()
}

type SingleTon struct {
	Attr string
}

func (this *SingleTon) TestFunc() {
	fmt.Println(this.Attr)
}

var (
	once     sync.Once
	instance *SingleTon
)

func GetInstance(str string) *SingleTon {
	once.Do(func() {
		instance = &SingleTon{Attr: str}
	})

	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			s := GetInstance("test:" + strconv.Itoa(i))
			s.TestFunc()
		}()
	}
	time.Sleep(1e5)
}
