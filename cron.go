// cron project main.go

package main



import (

	"fmt"

	"io/ioutil"

	"net/http"

	"time"
	"os"

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

	fmt.Print("/")

	return

}



func main() {
	param := os.Args
	data := make( map[string]string)
	len := len(param)-1
	// fmt.Println(len)
	for i=1;i<len;i+=2{
		// fmt.Println(param[i])
		data[param[i]] = param[i+1]
	}
	//default	
	if data["-t"]<="0" {
		data["-t"] = "5"
	}
	if data["-url"] == ""{
		fmt.Println("Url is empty!");
		return;
	}
	fmt.Println(data["bbb"])
	i = 0	
	fmt.Println("Start..")
	for {

		go request_cron(data["-url"])

		time.Sleep(4 * 1e9)

	}



}


