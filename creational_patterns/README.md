## 1. 简单工厂模式

类图：
![](http://img.blog.csdn.net/20170826202731964)

和传统的Java等面向对象语言的继承机制不同，在Go语言中并不提倡使用继承，但是在某些场景下要用到继承的话，可以用下面的方式实现继承关系。

```
type Animal struct {
  Category string
}

type Dog struct {
  Animal
  Name string
}
```
简单工厂模式实现起来也比较简单，主要代码如下：

```
type AnimalFactory struct {
}

func NewAnimalFactory() *AnimalFactory {
	return &AnimalFactory{}
}

func (this *AnimalFactory) CreateAnimal(name string) Action {
	switch name {
	case "bird":
		return &Bird{}
	case "fish":
		return &Fish{}
	case "dog":
		return &Dog{}
	default:
		panic("error animal type")
		return nil
	}
}
```

## 2. 工厂方法模式

简单工厂模式是通过传递不同的参数生成不同的实例，缺点就是扩展不同的类别时需要修改代码。

工厂方法模式为每一个product提供一个工程类，通过不同工厂创建不同实例。

类图：

![func](http://p0.qhimg.com/t0136c0b7b213d65023.gif)

实现过程：

- 1. 工厂方法定义一个创建struct的接口，让子struct去实现。

```
type AnimalFactory interface {
        CreateAnimal() Action
}

```

- 2. BirdFactory创建一个Bird的实例

```
type BirdFactory struct {
}

func (this *BirdFactory) CreateAnimal() Action {
        return &Bird{}
}

```

- 3. 工厂方法使用

```
    bFactory := &BirdFactory{}
    bird := bFactory.CreateAnimal()
    bird.Move(100)

```

### 3. 建造者模式

建造者模式将一个复杂的对象与它的表示分离，同样的创造过程可以建造出不停的表示。比如汽车，它包括商标、车轮、颜色、发送机等各种部分。而对于大多数用户而言，无须知道这些部件的装配细节，也几乎不会使用单独某个部件，而是使用一辆完整的汽车，可以通过建造者模式对其进行设计与描述，建造者模式可以将部件和其组装过程分开，一步一步创建一个复杂的对象。

类图：
![builder](http://img.blog.csdn.net/20170903163110850)

Go语言实现：

```
package main

import (
	"fmt"
)

//产品角色
type Car struct {
	Brand string
	Type  string
	Color string
}

func (this *Car) Drive() error {
	fmt.Printf("A %s %s %s car is running on the road!\n", this.Color, this.Type, this.Brand)
	return nil
}

//建造者角色
type Builder interface {
	PaintColor(color string) Builder
	AddBrand(brand string) Builder
	SetType(t string) Builder
	Build() Car
}

//具体的建造者
type ConcreteBuilder struct {
	ACar Car
}

func (this *ConcreteBuilder) PaintColor(cor string) Builder {
	this.ACar.Color = cor
	return this
}

func (this *ConcreteBuilder) AddBrand(bnd string) Builder {
	this.ACar.Brand = bnd
	return this
}

func (this *ConcreteBuilder) SetType(t string) Builder {
	this.ACar.Type = t
	return this
}

func (this *ConcreteBuilder) Build() Car {
	return this.ACar
}

//导演着角色
type Director struct {
	Builder Builder
}

func main() {
	dr := Director{&ConcreteBuilder{}}
	adCar := dr.Builder.SetType("SUV").AddBrand("奥迪").PaintColor("white").Build()
	adCar.Drive()

	bwCar := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bwCar.Drive()

}
```

## 4. 单例模式

单例模式，顾名思义就是在程序的运行中只产生一个实力。在Go实现上也有多种形式。

1. 懒汉模式.
  这种方式实现起来特别简单，直接判断一个实力是不是为`nil`， 如果是，则新生成；否则返回已有的。但它和多数语言一样，只适合用在单线程。

```
type SingleTon struct {
}

var instance *SingleTon

func GetInstance() *SingleTon {
	if Instance == nil {
		instance = &SingleTon{}
	}

	return instance
}
```

2. 使用加锁机制

在Go语言中有个基础对象`sync.Mutex`，可以实现协程之间的同步逻辑。

```
var mu sync.Mutex

func GetInstance() *SingleTon {
  mu.Lock()
  defer mu.Unock()

	if Instance == nil {
		instance = &SingleTon{}
	}

	return instance
}
```

3. `sync.Once`用法

在Go中还有一个更简洁的方法就是使用`sync.Once`，它可以在多协程中起到控制作用。实现起来也非常简单。

```
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
```
测试代码如下，从运行结果来看，都是一致的。

```
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			s := GetInstance("test:" + strconv.Itoa(i))
			s.TestFunc()
		}()
	}
	time.Sleep(1e5)
}

```
