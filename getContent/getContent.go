package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
  根据提供的url 获取返回信息内容
 */
func GetContents(url string) (string ,error) {

	resp,err := http.Get(url)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get content failed status code is %d ",resp.StatusCode)
	}

	bytes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "" , nil
	}
	return string(bytes),nil
}

func main() {
	url := "https://hz.zu.anjuke.com/"
	contents,err := GetContents(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(contents)
}
