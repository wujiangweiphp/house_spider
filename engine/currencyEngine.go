package engine

import (
	"house_spider/fetch"
	"log"
)

type Scheduler interface {
    Submit(Request)
    WorkChan() chan Request
    Run()
	ReadyNotify
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}

type CurrencyEngine struct {
	Sch Scheduler
	WorkerNum int
}



func (e *CurrencyEngine)Run(seeds ...Request){
	out := make(chan RequestResult)
	e.Sch.Run()

	for i := 0;i<e.WorkerNum; i++ {
		createWorker(e.Sch.WorkChan(),out,e.Sch)
	}

	//1. 将启动url加入请求队列
	for _,r := range seeds {
		e.Sch.Submit(r)
	}

	//itemCount := 0
	for {
		result := <- out
		//2.循环发送请求
		/*
		for _, item := range result.Items.(map[string]interface{}) {
			log.Printf("%d %v \n",itemCount,item)
			itemCount++
		}*/

		for _,r := range result.R {
			e.Sch.Submit(r)
		}
	}

}

func createWorker(in chan Request,out chan RequestResult,ready ReadyNotify) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (RequestResult,error)  {
	//2.1 抓取网页
	contents, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url error url is %s ,error is  %v\n", r.Url, err)
		return RequestResult{},err
	}
	//2.2 解析
	result := r.ParserFunc(contents)

	return result,nil
}
