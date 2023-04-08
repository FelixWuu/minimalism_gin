# 启动一个 Gin 服务

Gin 是一个用 go 语言编写的 web 框架，它具有高性能、低内存占用、简洁易用等特点，被广泛应用于构建 restful api 和微服务。

Gin 框架提供了丰富的功能和中间件支持，包括路由、参数解析、验证器、日志、错误处理、静态文件服务、模板引擎等。它还支持多种 http 方法和自定义路由规则，方便开发者根据实际需求进行定制化开发。此外，Gin 框架还提供了优雅的错误处理机制，通过捕获 panic 异常并输出详细的错误信息，让开发者更加方便地定位和排查问题。

相较于其他 web 框架，Gin 框架不仅拥有高效的性能表现，而且使用起来也非常简单，学习门槛较低，可以大大提升开发效率。现在，我们来看一段代码，开始学习 Gin 的使用:

> 你可以在我的 github 仓库 [minimalism_gin ](https://github.com/FelixWuu/minimalism_gin) 上看到本系列的代码和文档
> 

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello NutCat!")
	})

	router.Run(":8080")
}
```

这段代码的含义

- `router := gin.Default()` : 创建了一个默认的路由
    - 路由（Routing）：是指将 HTTP 请求与相应的处理程序相关联的机制。在 Web 应用程序中，客户端通过 HTTP 协议向服务器发送请求以获取数据或执行某些操作。路由就决定了当接收到该请求时，应该由哪个处理程序来处理。
    - 通俗易懂的讲，比如微博
      
        ![https://pic.imgdb.cn/item/643190b30d2dde57770f5296.png](https://pic.imgdb.cn/item/643190b30d2dde57770f5296.png)
        
        热门、同城、榜单、登录等按钮实际上都是路由，点击它会调用不同的处理程序，返回特定的一些结果。如登录按钮，点击后会弹出登录界面，让用户填写账户密码进行登录。
    
- `router.GET()` : 指定请求方式，这里指定为 GET 方式，路由规则为 /index，并有一个路由函数，处理了某些事情。
- `router.Run()`：启动服务并监听 ip，":8080" 指定了端口
    - 在 Gin 中有另一种启动方式，代码如下：`http.ListenAndServe(":8080", router)`
    - 这种方式使用的是原生 HTTP 服务的方式启动，效果与 `router.Run()` 相同
        - `router.Run()` 本质上是 `http.ListenAndServe` 的进一步封装

现在我们已经启动了一个 Gin 服务了，我们现在在浏览器中访问 [http://127.0.0.1:8080/index](http://127.0.0.1:8080/index) 即可访问到我们的服务了

![https://pic.imgdb.cn/item/643193df0d2dde57771448c9.png](https://pic.imgdb.cn/item/643193df0d2dde57771448c9.png)