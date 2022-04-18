package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err) // Print()
	}

	defer resp.Body.Close() // 延时关闭

	body, err := ioutil.ReadAll(resp.Body) // IO工具读取数据，返回字节数组

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body)) // 字节数组转字符串输出
}
