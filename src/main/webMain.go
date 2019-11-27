package main

//http 需要引入  net/http包
import (
	"fmt"
	"net/http"
)

//http.Request前面为什么要加* 号
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello web")
	fmt.Fprint(w, "欢迎来到go世界")
}

func main() {
	//设置访问的路由
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("发生致命错误")
	}

}
