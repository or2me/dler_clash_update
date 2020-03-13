package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var head string = "port: 7890\nsocks-port: 7891\nmode: Rule\nlog-level: info"

func main() {
	listUrl := "https://dler.cloud/subscribe/7STNYcUTWCmuEMiA?protocols=ss&list=clash"

	// Get the data
	resp, err := http.Get(listUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var datas string
	datas = strings.Replace(string(data), "proxies", "Proxy", 1)
	datas = strings.Replace(datas, "---", "", 1)
	datas = strings.Replace(datas, "...", "", 1)

	// ioutil.WriteFile("img1.jpg", data, 0644)
	var s []string
	s = strings.Split(datas, "\n")
	// for i, value := range s {
	// 	if (i-2)%7 == 0 {
	// 		head += "  " + value + "\n"
	// 	}
	// }
	var buffer bytes.Buffer
	for i := 1; i < len(s)-1; i++ {
		b := strings.Contains(s[i], "name:")
		if b {

			buffer.WriteString("  ")
			buffer.WriteString(strings.Replace(s[i], "name: ", "", 1))
			buffer.WriteString("\n")
		}
	}
	//添加节点
	head += datas
	//添加所有节点名称到代理组
	head += "Proxy Group:\n- name: load-balance\n  type: load-balance\n  proxies:\n" + buffer.String() + "\n  url: http://www.gstatic.com/generate_204\n  interval: \"600\""
	//添加rule
	head += "\nRule:\n  - 'MATCH,load-balance'\n"
	writeToFile2(head)
	fmt.Println("update finished")
}
func writeToFile2(msg string) {
	if err := ioutil.WriteFile("config.yaml", []byte(msg), 7777); err != nil {

		fmt.Println(err.Error())
	}
}
