package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	//http.HandleFunc("/test", rootHandler)
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//打印环境变量
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
	}
	//多路复用处理函数
	mux := http.NewServeMux()
	//debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

//健康检查路由
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
	//w.Header().Set()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("uid")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	//当前系统的环境变量的version配置
	//设置环境变量默认值
	ver := getEnvDefault("VERSION", "1.2.3")
	io.WriteString(w, fmt.Sprintf("VERSION : %s\n", ver))
	//03.客户端IP，httpcode
	cIP := ClientIP(r)
	log.Printf("Bingo!Response code:%s\n", "200")
	log.Printf("Bingo!The client IP is:%s\n", cIP)
}

func ClientIP(r *http.Request) string {
	xForwardFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

//设置默认环境变量
func getEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	os.Setenv(key, defVal)
	return defVal
}
