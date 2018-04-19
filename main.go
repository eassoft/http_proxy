package main

import (
	"hm/eclibs/logger"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/eassoft/http_proxy/config"
)

/*
根据子目录添加代理
*/
func AddProxy(proxy *config.ProxyConfig) {
	logger.Debug("proxy ", proxy.Pattern+" "+proxy.Scheme+" ", proxy.Host)
	http.HandleFunc(proxy.Pattern, func(w http.ResponseWriter, r *http.Request) {
		director := func(req *http.Request) {
			req = r
			req.URL.Scheme = proxy.Scheme
			req.URL.Host = proxy.Host

		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})
}

func main() {
	//根据配置文件添加
	for _, proxy := range config.ProxyConfigArr {

		AddProxy(proxy)
	}

	//log.Fatal(http.ListenAndServe(":8001", nil))
	logger.Debug("ListenAndServeTLS :443")
	log.Fatal(http.ListenAndServeTLS(":443", "conf/server1.pem", "conf/server1.key", nil))
}
