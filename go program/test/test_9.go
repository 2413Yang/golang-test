package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//发起GET请求
	resp, err := http.Get("https://comment.bilibili.com/xxxx.xml")
	if err != nil {
		fmt.Println("请求失败: %s",err)
		return
	}
	defer resp.Body.Close()

	//读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:%s",err)
		return
	}

	//打印页面内容
	fmt.Println(string(body))
}