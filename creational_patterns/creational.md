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
