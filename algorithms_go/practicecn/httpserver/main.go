package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
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
	//客户端IP，httpcode
	fmt.Printf("remoteAdd is:%s\n", r.RemoteAddr)
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
