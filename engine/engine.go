package engine

import (
	"house_spider/fetch"
	"log"
)

type Engine struct {
	WorkerNum int
	Ch chan Request
}

func (e *Engine)Run(seeds ...Request){
	var requests []Request

	//1. 将启动url加入请求队列
	for _,r := range seeds {
		requests = append(requests,r)
	}

	for {
		//2.循环发送请求
		for _, r := range requests {

			//2.0 出队列
			requests = requests[1:]
			//2.1 抓取网页
			contents, err := fetch.Fetch(r.Url)
			if err != nil {
				log.Printf("fetch url error url is %s ,error is  %v\n", r.Url, err)
				continue
			}
			if contents == "" {
				log.Printf("contents error url is %s ,error is  %v\n", r.Url, err)
				continue
			}
			//2.2 解析
			result := r.ParserFunc(contents)

			// 2.3 如果解析出请求 再次追加给 requests
			if len(result.R) > 0 {
				requests = append(requests, result.R...)
			}
		}
	}

}
