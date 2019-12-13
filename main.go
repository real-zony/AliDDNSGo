package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var commandModel CommandModel

func main() {
	initCommandModel()
	loadConfig()

	fmt.Println(getPublicIp())
}

func initCommandModel() {
	commandModel.FilePath = flag.String("f", "", "指定自定义的配置文件，请传入配置文件的路径。")
	commandModel.Interval = flag.Int("i", 0, "指定程序的自动检测周期，单位是秒。")

	flag.Parse()
}

func loadConfig() {
	if *commandModel.FilePath == "" {
		dir, _ := os.Getwd()
		fmt.Printf("%s\n", dir)
	}
}

func getPublicIp() string {
	resp, err := http.Get("http://members.3322.org/dyndns/getip")
	if err != nil {
		log.Printf("获取公网 IP 出现错误，错误信息：%s", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return string(bytes)
}
