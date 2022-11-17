package module

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type RProxy struct {
	Remote *url.URL
}

func GoReverseProxy(this *RProxy) *httputil.ReverseProxy {
	remote := this.Remote
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(request *http.Request) {
		//targetQuery := remote.RawQuery
		request.URL.Scheme = remote.Scheme
		request.URL.Host = remote.Host
		request.Host = remote.Host
		request.URL.Path, request.URL.RawPath = joinURLPath(remote, request.URL)
		request.Header.Set("X-Real-Ip", "47.122.2.224")
		request.Header.Set("X-Forwarded-For", "47.122.2.224")
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("User-Agent", "curl/7.29.0")
		request.Header.Set("Refer", "47.122.2.224")
		request.Header.Set("Accept", "*/*")
		request.Method = "POST"

		fmt.Println(request.RemoteAddr)
		//str, _ := ioutil.ReadAll(request.Body)

		//if targetQuery == "" || request.URL.RawQuery == "" {
		//	request.URL.RawQuery = targetQuery + request.URL.RawQuery
		//} else {
		//	request.URL.RawQuery = targetQuery + "&" + request.URL.RawQuery
		//}
		//if _, ok := request.Header["User-Agent"]; !ok {
		//	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
		//}
		log.Println("request.URL.Path：", request.URL.Path, "request.URL.RawQuery：", request.URL.RawQuery,
			"request-body : ")
	}

	// 修改响应头
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Add("Access-Control-Allow-Origin", "*")
		response.Header.Add("Reverse-Proxy-Server-PowerBy", "https://qianguopai.com")

		return nil
	}

	return proxy
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
