// cron project main.go

package main

import (
	"fmt"

	"io/ioutil"

	"net/http"

	"os"
	"strconv"
	"time"
)

//请求总次数
var total int

//是否输出结果
var result bool

func request_cron(url string) {

	total++

	fmt.Println("Request num :", total)

	resp, err1 := http.Get(url)

	if err1 != nil {

		fmt.Print("Err1")

		return

	}

	defer resp.Body.Close()

	data, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {

		fmt.Print("Err2")

		return

	}

	if result {
		fmt.Print(string(data))
	}
	return

}

//type MyFloat64 float64

func main() {
	param := os.Args
	data := make(map[string]string)
	len := len(param)
	for i := 1; i < len; i += 2 {
		// fmt.Println(param[i])
		if (i+1) < len && param[i+1] != "" {
			data[param[i]] = param[i+1]
		} else {
			data[param[i]] = ""
		}
	}
	// fmt.Println(data)
	_, ok := data["-help"]
	_, ok2 := data["-h"]
	if ok || ok2 {
		fmt.Println("cron help!")
		fmt.Println("-url        :请求的url地址")
		fmt.Println("-t            :每次请求的间隔秒  5")
		fmt.Println("-n            :每次请求多少次   1")
		fmt.Println("-r            :是否输出结果   true")
		return
	}
	//default
	//请求的url地址
	if data["-url"] == "" {
		fmt.Println("Url is empty!")
		return
	}
	//每次请求的间隔秒
	if data["-t"] == "" {
		data["-t"] = "5"
	}
	//每次请求多少次
	if data["-n"] == "" {
		data["-n"] = "1"
	}
	//是否输出结果
	if data["-r"] == "" {
		data["-r"] = "true"
	}
	if data["-r"] == "true" {
		result = true
	} else {
		result = false
	}
	//fmt.Println(data["bbb"])
	total = 0
	fmt.Println("Start..")
	//转化为float64位
	second_h, _ := strconv.ParseFloat(data["-t"], 64)
	//转化为微秒
	second_m := second_h * 1e9
	//转化为符合time.Sleep的参数
	second_num := time.Duration(second_m)
	n, _ := strconv.ParseInt(data["-n"], 10, 64)
	var i int64
	for {
		fmt.Println("Next..")
		for i = 0; i < n; i++ {
			go request_cron(data["-url"])
		}

		time.Sleep(second_num)

	}

}
