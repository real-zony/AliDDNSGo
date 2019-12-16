package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var commandModel CommandModel
var configModel ConfigurationModel

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
	var configFile string
	if *commandModel.FilePath == "" {
		dir, _ := os.Getwd()
		configFile = path.Join(dir, "settings.json")
	} else {
		configFile = *commandModel.FilePath
	}

	// 打开配置文件，并进行反序列化。
	f, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("无法打开文件：%s", err)
		os.Exit(-1)
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)

	if err := json.Unmarshal(data, &configModel); err != nil {
		log.Fatalf("数据反序列化失败：%s", err)
		os.Exit(-1)
	}
}

func getPublicIp() string {
	resp, err := http.Get(GetPublicIpUrl)
	if err != nil {
		log.Printf("获取公网 IP 出现错误，错误信息：%s", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return string(bytes)
}

func getSubDomains() {
	client, err := alidns.NewClientWithAccessKey("", configModel.AccessId, configModel.AccessKey)
	request := alidns.CreateDescribeDomainRecordsRequest()

	request.Scheme = "https"
	request.Domain = configModel.MainDomain

	_, err = client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
