# [译]利用函数式参数构建友好的 API

> 原文 [Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)  
> 作者: Dave Cheney  
> 译者: Sam Fu

我想从一个故事说起。

时间到了 2014 年年底（这是本文演讲当年），你的公司正在推行一个革命性的分布式社交网络项目，你们项目很明智地选择了 go 作为产品的开发语言。

你被分配到的任务是负责编写一个关键的服务组件，这大概看起来像是这样:

```go
package gplusplus

import "net"

type Server struct {
    listener net.Listener
}

func (s *Server) Addr() net.Addr
func (s *Server) Shutdown()

// NewServer returns a Server listen on Addr
func NewServer(addr string) (*Server, error) {
    l, err :=  net.Listen("tcp", addr)
    if err != nil {
        return nil, err
    }
    srv := Server{ listener: l }
    go srv.run()
    return &srv, nil
}

```

我们有一些还没导出的变量需要进行初始化，同时我们需要启动一个 go routine 来接收服务的请求。

这个包只提供了一个 api，这很容易上手使用。

但是，这里有个问题，当你第一个 beta 发布后，新功能的需求开始不断的涌现。

移动端经常产生慢响应或者扎堆的停止响应，你需要给这些慢连接提供断开连接的功能。

在安全性能日益重视的氛围下，你的缺陷追踪板开始被支持安全连接的需求所填满。

然后你获悉某些用户将你的 Server 在很小型的 VPS 中运行，他们需要一种方式去限制最大同时在线的客户端。

接下来就是来自一班被机器攻击的用户所需要的限制并发请求功能。

还有 blabla....

然后，基于上面提及的我们开始动手进行升级

```go
// NewServer returns a Server listen on Addr
// clientTimeout defines the maximum length of an idle connection, or forever if not provided.
// maxconns limits the number of concurrent connections.
// maxconcurrent limits the number of concurrent connections from a single IP address.
// cert is the TLS certificate for the connection.
func NewServer(addr string, clientTimeout time.Duration, maxconns int, maxconcurrent int, cert *tls.Cert) (*Server, error)
```

当这页 ppt 已经无法轻松容纳这个函数声明的时候，这是个不好的预兆。

举手看看，你们谁写过这样的 api？

谁曾因依赖这个一样的 api 而导致破坏代码的？

很明显，这个 api 不单累赘，而且脆弱，而且这很不直观。

使用者一看这个 api 并不能在短时间内了解到那些参数是必须的，那些是可选的。

举个例子，如果我只是想拿这个 Server 来做一个测试，那么我怎么知道是否需要一个真实可用的 TLS 证书，如果我不提供，那我该如何去传这个参数。

如果我不需要 maxconns, maxconcurrent 这些参数那我需要提供什么参数去调用？我可以用 0 吗？用 0 传进去听起来可行，但是，如果传 0 进去这取决于内部的实现逻辑，这可能导致你限制了最大总并发连接数为 0 （这并不是你想要的，你只想单纯忽略这个参数）。

在我看来，编写这样的 api 其实很容易，只要你让调用者有责任地正确使用它。

虽然这个例子可以看做是故意夸张构造并且文档单薄的复杂例子，但我相信它反映出的是华丽而脆弱的真实的 api 问题。

好了，我定义好问题本身了，我们来看看一些解决方案：

```go
// 将工作细分到多个函数
NewServer(addr string)(*Server, error)
NewTLSServer(addr string, cert *tls.Cert)(*Server, error)
NewServerWithTimeout(addr string, timeout time.Duration) (*Server, error)
NewTLSServerWithTimeout(addr string, cert *tls.Cert, timeout time.Duration) (*Server, error)
```

相对于将所有参数编排进一个函数中的方法，一种方式是将这些组合分散到一组函数中，按需调用。

使用这种方式后，需要 TLS 就调用 TLS 的方法，需要最大超时的就使用包含最大超时的方法。

很不幸地，如你所见如果把所有组合都列出来，一下子很快就会被函数堆所淹没。

让我们再来尝试另一种方法使 api 可配置：

```go
// 使用struct进行配置
type Config struct {
	timeout time.Duration
	Cert *tls.Cert
}
func NewServer(addr string, config Config)(*Server, error)
```

使用 struct 来传递这是一种比较普遍的方式，这个方法有一些优点：

能保证 api 函数签名不变的前提下，随着时间推移去不断增加 Config 内部的属性；

这种方式可以带来更好的文档，原本需要在函数上方添加大量的参数注释，现在只需要在结构化的 struct 旁边进行标注就行了；

潜在地，这也允许调用者去使用零值去表示他们需要指定配置项， 并让它们按照默认行为来表现。

然而这也会导致一些问题，当某些属性的零值是有明显意义的：

```go
type Config struct {
	Port int
}

func main () {
	srv, _ := NewServer("localhost", Config{
	Port: 0, // 这里port填0会有问题
})
}
```

如上面代码所示， 如果填 0 的话，会返回一个 8080 口的 \*Server。这样的话你将没办法显式传数值 0 企图表示你不想传这个参数，然后让系统去自动分配端口，因为无法分辨 0 是被显式赋值还是说是本身的零值。

```go
// 我就是想要一个Server，别给我整些有的没的
func NewServer(addr string, config Config) (*Server, error)

func main() {
	srv, _ := NewServer(addr, Config{}) // 因为各种理由传空配置
}
```

很多时候调用者就是需要 api 的默认行为，即使他们并不需要去改动配置的参数，但是，他们还是不得不去传第二个参数。因此别人尝试去阅读你代码的时候，看到这个神奇的空配置参数，他们会因此而懵逼，你传了参数却是空的，也不懂有没有什么隐含的行为。

我觉得呢，这个是不对的。

为什么别人要给你的 api 去构造一个空的参数呢？就是为了去满足函数签名吗？

```go
// maybe adding indirection will help
func NewServer(addr string, config *Config) (*Server, error)
func main() {
	srv, _ := NewServer("localhost", nil) // 支持默认参数
	config := Config{ Port: 9000 }
	srv2, _ := NewServer("localhost", &config)
	config.Port = 9001 // 这时会发生什么
}
```

一种常用的方式去传递空数值就是使用指针传递，那么调用者就可以使用 nil 表示空值而不是使用构造形式构造空值。

在我看来，这样还是保持跟前面一样的问题，还加了一点问题：

虽然我们还是需要满足函数签名而给上一个 nil， 但是一个 nil 就可以代表使用默认的参数。

这引申出一个问题，到底直接传 nil 跟传一个空值的指针有区别吗？ (nil vs &Config{})

更令包作者和调用者疑惑的是，现在 Server 和调用者可以有同一份 config 的引用，这又会带来一个问题，要是这个 config 在 NewServer 调用后修改那会怎样？

我相信一个写得好的 api 是不需要调用者去创建空的值只是去为了满足罕有的用例。
我深信我们作为 go 程序员，应该努力去确保 nil 不成为一个公共方法必须传递的参数。

而当我们真的需要去传递配置信息的时候，它应该是具备自述性，尽可能具有表现力的。

综上几点，我想讨论一些我认为更好的做法：

```go
// 可变配置
func NewServer(addr string, config ...Config)(*Server, error)
func main() {
	srv, _ := NewServer("localhost") // 默认参数
	srv2, _ := NewServer("localhost", Config{
		Timeout: 300 * time.Second,
		MaxConns: 10,
	})
}
```

为了移除强制但未经常使用 config 参数的问题，我们可以修改 NewServer 函数去支持可变数量参数。

与传递 nil 和空值不同，作为默认参数的标志，利用可变参数的特性根本不用传入任何参数。

这在我的书中，它解决了 2 个大问题：

首先，调用默认行为的方式做到了尽可能的简洁；

其次，NewServer 只接受 Config 类型的配置，不可以使用指针类型，排除了使用 nil 传递的情况，并且保证调用者不会获取一份配置引用（从而发生调用后修改参数的情况）。

我觉得这样已经很好了。

但如果我们深究下去，这还是有很多问题的。

很明显我们期望的只是接收最多一个参数，但是使用可变参数的形式，这个实现不得不去处理多个传入参数，甚至可能是矛盾的配置结构体。

那有没有更好的办法去使用可变参数来达到更好表现力的，按需取参的效果呢？

我认为就是这样：

```go
func NewServer(addr string, options ...func(*Server)) (*Server, error)
func main() {
	srv, _ = NewServer("localhost")

	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}

	tls := func(srv *Server) {
		config := tls.loadTLSConfig()
		srv.listener = tls.NewListener(srv.listener, &config)
	}

	srv2, _ := NewServer("localhost", timeout, tls)
}
```

在这我要说明的是这个想法来自 Rob Pike 的一篇博客[Self referential functions and design]([Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)), 今年(2014)1 月的文章， 我鼓励大家都去看看。

跟前面例子不同的关键点是，对 Server 的可配置项并不是包含在 struct 里的变量，而是可以操作 Server 本身的函数。

如果前面的例子一样，可变参数给了我们需要默认行为时的紧凑表达。

而当需要配置项的时候，我就传入可对 Server 本身操作的函数。

timeout 字段就是简单的把 Server 中的 timeout 修改

而 tls 就有点复杂，它把原有的 Server 用 tls listener 又包了一层，从而转换成安全的 listener。

```go
func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{ listener: l }
	for _, option := range {
		option(&srv)
	}
	return &srv, nil
}
```

在实现内部，执行的 option 的方式很直观。

我们在定义完 Server 后将它的地址作为 option 的参数。很明显，如果 options 数量为 0 ，那么应用 option 的 loop 将不会循环。

使用这种方式，我们可以使 api 拥有如下特性:

- 明智的默认调用方式
- 高度可配置
- 经历得了开发周期考验
- 自成文档
- 新手友好
- 不用去找空值和 nil 来讨好编译器
