# 实验要求

编程 web 应用程序 cloudgo-io。 请在项目 README.MD 给出完成任务的证据！


# 实验结果


## 1. 支持静态文件服务

访问http://localhost:9090/static/

进入静态服务

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fxbbtn2hbkj20p506y749.jpg)

![image](https://wx4.sinaimg.cn/mw690/b41a0581gy1fxbbtn2eacj20mi03e3ye.jpg)

## 2. 支持简单 js 访问


访问http://localhost:9090/

此时上面的Welcome是通过服务器的js 访问得到的

```
func apiTestHandler(formatter *render.Render) http.HandlerFunc {

		return func(w http.ResponseWriter, req *http.Request) {
			formatter.JSON(w, http.StatusOK, struct {
				WELCOME      string `json:"WELCOME"`
				
			}{WELCOME: "Welcome"})
	}
}

```

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fxbbtn37lvj20me041q2q.jpg)


## 3.提交表单，并输出一个表格

输入username 和 password， 将返回输入内容的表单

![image](https://wx3.sinaimg.cn/mw690/b41a0581gy1fxbbtn2ytgj20l707g74h.jpg)




## 4.对 /unknown 给出开发中的提示，返回501

访问127.0.0.1:8080/api/unknown

![image](https://wx4.sinaimg.cn/mw690/b41a0581gy1fxbbtn8w6kj20mt03xdfr.jpg)