package parser

import (
	"fmt"
	"house_spider/engine"
	"regexp"
)

var (
	/*
	  <a href="https://hz.zu.anjuke.com/fangyuan/xihu/" title="西湖租房"
                                       class="">西湖</a>
	 */
	ParseArea = regexp.MustCompile(`<a href="(https://hz.zu.anjuke.com/fangyuan/[a-zA-Z]+/)"[^>]+>([^<]+)</a>`)
)

//解析区域
func ParserArea(contents string)engine.RequestResult{
	result := ParseArea.FindAllStringSubmatch(contents,-1)
	var res engine.RequestResult
	for _,r := range result {
		if r[1] != "" {
			res.R = append(res.R,
				engine.Request{
					Url:r[1],
					ParserFunc: ParserList,
			    })
		}
		fmt.Printf("区域：%s ,link : %s \n",r[2],r[1])
	}
	return res
}

