package main

import (
	"house_spider/engine"
	"house_spider/house/parser"
	"house_spider/scheduler"
)

func main() {
	e := &engine.CurrencyEngine{
		Sch:&scheduler.QueueScheduler{},
		WorkerNum:5,
	}
	e.Run(engine.Request{
		Url:"https://hz.zu.anjuke.com/",
		ParserFunc:parser.ParserArea,
	})
}
