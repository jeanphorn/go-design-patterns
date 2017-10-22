# 1. 简述
-------------------
在面向对象的编程语言中（如java，C++）设计模式的概念广为人知, 应用的也非常广泛。设计模式让我们的代码变得灵活起来，具有很强的扩展性。但在与C语言比肩的Go语言中，设计模式的概念并没有十分突出，甚至很少听到。在Go的开发中，借鉴design pattern的理念同样回味无穷我们的开发带来极大的便利。

![go-logo](http://img.blog.csdn.net/20170826191635943)

# 2. Design pattern in go 学习实践
-------------------
老生常谈，还是从设计模式的三大类：创建型模式（Creational Patterns）、结构型模式（Structural Patterns）、行为型模式（Behavioral Patterns）开始介绍。

## 2.1 创建型模式
创建型设计模式提供了一种在创建对象的同时隐藏创建逻辑的方式，使得程序在判断针对某个给定实例需要创建哪些对象时更加灵活。

|模式|描述|
|:--:|:--:|
|[简单工厂模式](https://github.com/jeanphorn/go-design-patterns/blob/master/creational_patterns/simplefactory.go)|提供一个接口，根据不同的输入生成不同的实例|
|[工厂方法模式](https://github.com/jeanphorn/go-design-patterns/blob/master/creational_patterns/funcfactory.go)|每个实例都有创建该实例的工厂|
|[建造者模式](https://github.com/jeanphorn/go-design-patterns/blob/master/creational_patterns/builder.go)|将一个复杂的对象与它的表示分离，同样的创造过程可以建造出不停的表示|
|[单例模式](https://github.com/jeanphorn/go-design-patterns/blob/master/creational_patterns/singleton.go)|在程序运行过程中之产生一个实例|

## 2.2 结构型模式

结构型模式是将对象或结构体通过一定方式组成一个更复杂的对象或结构体，使其具有更复杂的功能，就像搭积木一样。结构型的模式同通常包括外观模式，桥接模式，适配器模式，装饰模式等。

|模式|描述|
|:--:|:--:|
|[外观模](https://github.com/jeanphorn/go-design-patterns/blob/master/structural_patterns/facade.go)|引入子系统统一调用入口|
|[装饰模式](https://github.com/jeanphorn/go-design-patterns/blob/master/structural_patterns/decorator.go)|动态扩展已存在对象的功能|
|[代理模式](https://github.com/jeanphorn/go-design-patterns/blob/master/structural_patterns/proxy.go)|动态扩展已存在对象的功能|
