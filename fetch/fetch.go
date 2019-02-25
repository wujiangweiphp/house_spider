package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var tick = time.Tick(time.Millisecond * 1000)

/**
  根据提供的url 获取返回信息内容
 */
func Fetch(url string) (string ,error) {
	<-tick
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
