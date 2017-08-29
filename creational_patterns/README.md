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
