package main

import (
	"house_spider/engine"
	"house_spider/house/parser"
)

func main() {
	e := &engine.Engine{}
	e.Run(engine.Request{
		Url:"https://hz.zu.anjuke.com/",
		ParserFunc:parser.ParserArea,
	})
}
