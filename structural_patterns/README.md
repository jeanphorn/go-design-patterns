## 结构型模式(structural pattern)

结构型模式是将对象或结构体通过一定方式组成一个更复杂的对象或结构体，使其具有更复杂的功能，就像搭积木一样。结构型的模式同通常包括外观模式，桥接模式，适配器模式，装饰模式等。

### 1. 外观模式

它为一套复杂的调度子系统提供一个统一的接入接口。外部所有对子系统的调用都通过这个外观角色进行统一调用，降低子系统与调用者之间的耦合度。

那当前比较热门的微服务来说，一套服务（比如说短视频服务）包括若干子服务,如图(a)，如：音乐服务，短视频服务，计数服务，推荐子服务等。客户端不同的请求会使用不同的子服务。客户端视频创作会请求音乐素材，视频上传服务；浏览推荐视频需要请求视频推荐服务和计数服务。这种情况下，我们可以引入一个统一的gateway层作为外观模式，统一管理入口, 如图(b)。

![before](http://img.blog.csdn.net/20170924155726971)

**图(a)**

![after](http://img.blog.csdn.net/20170924155738037)

**图(b)**

外观模式类图：

![](https://www.packtpub.com/sites/default/files/Article-Images/B05180_01.png)

Go语言实现：

```
type Facade struct {
	M Music
	V Video
	C Count
}

func (this *Facade) GetRecommandVideos() error {
	this.V.GetVideos()
	this.C.GetCountByID(111)

	return nil
}

type Music struct {
}

func (this *Music) GetMusic() error {
	fmt.Println("get music material")
	// logic code here
	return nil
}

type Video struct {
	vid int64
}

func (this *Video) GetVideos() error {
	fmt.Println("get videos1")
	return nil
}

type Count struct {
	PraiseCnt  int64 //点赞数
	CommentCnt int64 //评论数
	CollectCnt int64 //收藏数
}

func (this *Count) GetCountByID(id int64) (*Count, error) {
	fmt.Println("get video counts")
	return this, nil
}

```

### 2. 装饰模式

装饰模式就是在不改变对象内部结构的情况下，动态扩展它的功能。它提供了灵活的方法来扩展对象的功能。


下面是一个简单的实现逻辑,通过Decorate来进一步装饰`Dressing`函数：

```
type Object func(string) string

func Decorate(fn Object) Object {
        return func(base string) string {

                    ret := fn(base)

                            ret = ret + " and Tshirt"
                                    return ret
                                        }
}

func Dressing(cloth string) string {
        return "dressing " + cloth
}

```

- 使用方式

```
f := Decorate(Dressing)
fmt.Println(f("shoes"))
```
