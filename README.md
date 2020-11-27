# V2ray RSS
---
### 说明
通过爬取免费v2ray链接分享，自动生成订阅链接

V2ray客户端推荐：[Qv2ray](https://github.com/Qv2ray/Qv2ray)

最近tg频道不更新了

(删仓了) ~~~请大家使用[PPS](https://github.com/xyzjhe/pps)这个优秀的项目吧~~~

### 更新日志
##### 2020/11/4
- 使用并发爬取多个url
- 修改install.sh
- 增加新版Docker镜像
- 增加自定义url列表
> 列表规则
```
# 修改 /etc/v2rss/v2list
加入订阅地址，每行一个，必须是完整网址(包含 http:// or https://)。
订阅地址放到上面，tg频道放到下面
# 启动参数加入 -c 例：
v2rss -p 5500 -c /etc/v2rss/v2list
```
##### 2020/11/2
- 聚合了2个订阅地址
- - [cdp2020](https://github.com/cdp2020/v2ray)
- - [freefq](https://github.com/freefq/free)
- 最好使用[cf worker](https://github.com/netnr/workers)直接反代上面的项目raw地址获取数据
##### 2020/10/18
- 添加Docker部署方式
- 修复空数据
##### 2020/10/16
- 修复str转byte含有空数据导致的base64编码失败
##### 2020/10/13
- 修复base64编码bug导致的数据返回异常
- 原CORS项目地址：[cors.zme.ink](https://github.com/netnr/workers)，请多多给些star鼓励作者。
- 使用新的CORS，CloudFlare Works部署，用于前端跨域， 项目地址：[cors-anywhere-cfworker](https://github.com/sunyuting83/cors-anywhere-cfworker)。希望同学们建立自己的cf worker并重新打包，分散压力。
##### 2020/10/8
- 添加Linux一键部署、卸载脚本（Nginx自行配置）

### Docker部署
#### 拉镜像
```
sudo docker pull v2rss/v2rss:beta
```
#### 起容器
```
# 默认启动
sudo docker run --name v2rss:beta --net=host -d v2rss/v2rss

# 自定义url列表启动
sudo docker run --name v2rss:beta --net=host -v /yourPath/configFile:/home/v2list -d v2rss/v2rss
```
> Docker部署默认监听端口5500

### Linux一键安装脚本
```
curl https://raw.githubusercontent.com/sunyuting83/v2rss-go/dev/install.sh |bash
```
> 一键脚本监听端口5500
### Linux一键卸载脚本
```
curl https://raw.githubusercontent.com/sunyuting83/v2rss-go/dev/uninstall.sh |bash
```

### 启动参数说明
| 参数  | 说明 | 默认值 |
| ------------ | ------------ | ------------ |
| -p | 监听端口号 | 3000 |
| -c | URL列表路径 | 无(使用程序内默认列表) |

#### 参数说明
浏览器访问 http://localhost:5500/?i=1&w=0&n=1

| 参数  | 说明 | 默认值 |
| ------------ | ------------ | ------------ |
| w | 启用代理访问（国内跨域访问） | 0(不开启) |
| n | 内容数字 | 1 |
| i | 合并数据数量 | 0(不合并) |

### nginx配置文件
```
server {
    listen 80;
    server_name yourdomain.com;
    location / {
        proxy_pass http://127.0.0.1:5500;
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
    access_log  your log path;
}
```
修改yourdomain.com成你的域名

修改your log path 成你的log文件路径

使用其他端口 请修改 proxy_pass 后面的端口号

- Win版使用说明
> 下载v2rss_x86_64.zip 并解压。双击运行。
在v2ray客户端订阅处填入订阅地址：
http://localhost:3000/?w=1&n=1
必须加w参数，用于国内访问。
其他参数请看说明文件
订阅成功后关闭窗口即可

##### 注：如使用其他端口，修改启动文件start.sh -p参数后面的端口号