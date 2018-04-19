# http_proxy
golang 写的反向代理服务器，按虚拟目录代理。

因为微信小程序必须用https,而在阿里提供的免费域名只能支持一个子域名。 需要把多个服务通过虚拟目录进行负载。于是谢了这个程序。自己的代码随便搞，灵活。比学习配置nginx 爽。
配置文件如下conf/config.json
[{
	"Pattern":"/", //虚拟目录
	"Scheme":"http", //协议
	"Host":"127.0.0.1:9900"//被代理的服务
	 
	 
},{
	"Pattern":"/qm/",
	"Scheme":"http",
	"Host":"127.0.0.1:9961"
},{
	"Pattern":"/img/sign/",
	"Scheme":"http",
	"Host":"127.0.0.1:9961"
}
]
 
