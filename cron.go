// cron project main.go

package main

import (
	"fmt"

	"io/ioutil"

	"net/http"

	"os"
	"time"
	"strconv"
)

var i int

func request_cron(url string) {

	i++

	fmt.Print(i)

	fmt.Print(":")

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

	fmt.Print(string(data))	

	return

}

//type MyFloat64 float64

func main() {
	param := os.Args
	data := make(map[string]string)
	len := len(param) - 1
	// fmt.Println(len)
	for i = 1; i < len; i += 2 {
		// fmt.Println(param[i])
		data[param[i]] = param[i+1]
	}
	//default
	//请求的url地址
	if data["-url"] == "" {
		fmt.Println("Url is empty!")
		return
	}
	//每次请求的间隔秒杀
	if data["-t"] == "" {
            data["-t"] = "5"
	}
	//每次请求多少次
	if data["-n"] == "" {
            data["-n"] = "1"
	}
	//fmt.Println(data["bbb"])
	i = 0
	fmt.Println("Start..")
//转化为float64位
	second_h,_ := strconv.ParseFloat(data["-t"],64)
//转化为微秒
	second_m := second_h*1e9
//转化为符合time.Sleep的参数
	second_num := time.Duration(second_m)
        n,_ := strconv.ParseInt(data["-n"],10,64)
	var i int64
	for {
		for i=0;i<n;i++{
     		    go request_cron(data["-url"])
		}
		fmt.Println("Next..")
		time.Sleep(second_num)

	}

}
