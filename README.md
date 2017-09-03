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
