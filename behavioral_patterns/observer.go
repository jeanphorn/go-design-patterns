package main

import (
	"fmt"
	"time"
)

type Event struct {
	Data string
}

type Observer interface {
	//更新事件
	Update(*Event)
}

// 被观察的对象接口
type Subject interface {
	//注册观察者
	Regist(Observer)
	//注销观察者
	Deregist(Observer)

	//通知观察者事件
	Notify(*Event)
}

type ConcreteObserver struct {
	Id int
}

func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.Id, e.Data)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Regist(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

func (cs *ConcreteSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}

	cobserver1 := &ConcreteObserver{1}
	cobserver2 := &ConcreteObserver{2}

	cs.Regist(cobserver1)
	cs.Regist(cobserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
