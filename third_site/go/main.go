package main

import (
	"flag"
	"fmt"
	"github.com/printspirit/gosdk"
)

var site *gosdk.ThirdApp

const (
	UID  = "third_test" //请修改为你在打印精灵上的账号和密码
	PASS = "third_test"
)

func main() {
	port := flag.Int("p", 8000, "WEB端口")
	ip := flag.String("ip", "", "SpiritCenter的安装IP地址")
	flag.Parse()

	if *ip == "" {
		site = gosdk.NewThirdApp("https://www.printspirit.cn", UID, PASS)
		
		fmt.Println("")
		fmt.Println("使用打印精灵官网作为标签编辑服务器")
		fmt.Println("使用SpiritCenter请用-ip设置IP地址")
		
	} else {
		// 对于SpiritCenter 用下面航替换
		center_url := fmt.Sprintf("http://%s:9011", *ip)
		site = gosdk.NewThirdApp(center_url, "", "")
	}

	Start(*port)
}
