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
	ParseList = regexp.MustCompile(`href="(https://hz.zu.anjuke.com/fangyuan/\d+)"\s*>([^<]+)</a>`)
)

//解析区域列表房源信息
func ParserList(contents string)engine.RequestResult{
	result := ParseList.FindAllStringSubmatch(contents,-1)
	var res engine.RequestResult
	for _,r := range result {
		if r[1] != "" {
			res.R = append(res.R,
				engine.Request{
					Url:r[1],
					ParserFunc:ParserHouse})
		}
		fmt.Printf("房源名称：%s ,link : %s \n",r[2],r[1])
	}
	return res
}

