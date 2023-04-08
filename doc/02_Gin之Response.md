# Response

> 你可以在我的 github 仓库 [minimalism_gin](https://github.com/FelixWuu/minimalism_gin) 上看到本系列的代码和文档
> 

## 1. 各式文件类型的响应

### 响应 String

我们定义一个函数，用于响应 String。

```go
func StringResponse(c *gin.Context) {
	c.String(http.StatusOK, "Hello, NutCat! This is your string")
}
```

上一节我们也看到了一个简单的响应例子，它在页面上返回了 Hello NutCat

![https://pic.imgdb.cn/item/643193df0d2dde57771448c9.png](https://pic.imgdb.cn/item/643193df0d2dde57771448c9.png)

同理，我们使用上节的方式，来响应 String

```go
package main

import (
	"github.com/FelixWuu/minimalism_gin/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/rsp/string", response.StringResponse)

	router.Run()
}
```

- `router.GET("/rsp/string", response.StringResponse)`： 我们直接传入函数给 router.GET 即可，实际上，main 函数里还是尽量不要出现业务逻辑
- `router.Run()`：我们没有像上一节一样填写 8080 端口了，这是应为 gin 默认就是访问 8080 端口

此时我们再次启动服务，访问 [http://127.0.0.1:8080/rsp/string](http://127.0.0.1:8080/rsp/string)，就可以看到我们输出的字符串了。

### 响应 JSON

```go
type JsonMsg struct {
	Name        string  `json:"name"`
	Score       float64 `json:"score"`
	Description string  `json:"description"`
}

func JSONResponse(c *gin.Context) {
	msg := JsonMsg{"NutCat", 99.9, "打工仔1号"}
	c.JSON(http.StatusOK, msg)
	c.JSON(http.StatusOK, gin.H{"Name": "FelixWuu", "Score": 100, "Description": "打工仔2号"})
}
```

- 我们可以直接传入一个 struct，它会被转换为 JSON 显示
- 除了 struct 外，我们还可以是用 `gin.H{}`
    - 在 Gin 框架中，gin.H 是一个类型为 `map[string]interface{}` 的数据结构，可以在处理 HTTP 请求时存储和检索键值对数据。下面是他的类型和官方的解释：
        - `type H map[string]any`
        - H is a shortcut for map[string]interface{}
    - gin.H 中的每个键都是一个字符串类型的字段名（key），而每个值（value）则是一个空接口类型，可以存储任何类型的数据。这使得 gin.H 非常灵活，并且可以用于许多不同的场景，如本小节的响应 JSON 数据。

在 main 函数中注册路由

```go
router.GET("/rsp/json", response.JSONResponse)
```

输入 [http://127.0.0.1:8080/rsp/json](http://127.0.0.1:8080/rsp/json)，这条路由在页面显示的内容如下：

```go
{"name":"NutCat","score":99.9,"description":"打工仔1号"}{"Description":"打工仔2号","Name":"FelixWuu","Score":100}
```

可以看到结构体是按原有顺序打印出来的，但 gin.H 并不是，原因就是它的底层是 map，是无序的。

### 响应 HTML

我们需要一个 HTML 文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>响应 HTML</title>
</head>
<body>
Hello NutCat~
</body>
</html>
```

然后我们可以写一个函数来响应 HTML 了

```go
func HTMLResponse(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
```

但事情并没有那么简单，首先 gin 并不知道 index.html 在哪里，这种写法显然无法让响应正常展示在浏览器上（实际上你可以试着启动一下，是会报错的）

对于 HTML，我们需要先要使用 `LoadHTMLGlob()`或者`LoadHTMLFiles()`方法来加载模板文件

```go
func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/rsp/string", response.StringResponse)
	router.GET("/rsp/json", response.JSONResponse)
	router.GET("rsp/html", response.HTMLResponse)

	router.Run()
}
```

现在输入 [http://127.0.0.1:8080/rsp/html](http://127.0.0.1:8080/rsp/html)，既可以看见响应出来的内容了。细心的你可能发现了我在 `HTMLResponse` 里写了一个空的 `gin.H{}` ，这是因为我们可以向 HTML 里面传参。

现在我们将 HTML 的 body 部分改为

```html
<body>
Hello NutCat~
现在的时间是 {{ .time}}
</body>
```

然后我们修改一下 `HTMLResponse` 让他可以穿参数给 HTML

```go
func HTMLResponse(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		// Pass in the current time
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}
```

启动服务访问  [http://127.0.0.1:8080/rsp/html](http://127.0.0.1:8080/rsp/html)，可以看到页面显示后端传入的当前时间 `Hello NutCat~ 现在的时间是 2023-04-09 00:57:25`

### 响应 XML

```go
funcXMLResponse(c *gin.Context) {
   c.XML(http.StatusOK, gin.H{"user": "NutCat", "score": 100.00, "level": 3})
}

```

```go
router.GET("rsp/xml", response.XMLResponse)
```

访问 [http://127.0.0.1:8080/rsp/xml](http://127.0.0.1:8080/rsp/xml)，得到：

```xml
<map>
<score>100</score>
<level>3</level>
<user>NutCat</user>
</map>
```

### 响应 YAML

```go
func YAMLResponse(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "FelixWuu", "score": 99.99, "level": 10})
}
```

```go
router.GET("rsp/yaml", response.YAMLResponse)
```

注：浏览器可能会跳转到下载连接，而不是显示出来，可以用 postman 看下结果

![https://pic.imgdb.cn/item/6431bd440d2dde57774e0da7.png](https://pic.imgdb.cn/item/6431bd440d2dde57774e0da7.png)