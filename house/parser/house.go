package parser

import (
	"fmt"
	"house_spider/engine"
	"house_spider/house/model"
	"regexp"
	"strconv"
)

var (
	BuildingNameParse = regexp.MustCompile(`<span class="type">小区：</span>\s+<a href="https://hangzhou.anjuke.com/community/view/\d+" class="link"  target="_blank" _soj="propview">([^<]+)</a>`)  //小区名称
	PayTypeParse      = regexp.MustCompile(`<span class="type">(付\d+押\d+)</span>`)  // 付3押1
	UnitTypeParse     = regexp.MustCompile(`<span class="type">户型：</span>\s+<span class="info">([^<]+)</span>`)  // 3室2厅2卫
	AreaParse         = regexp.MustCompile(`<span class="type">面积：</span>\s+<span class="info">(\d+)平方米</span>`)     //面积
	TowordParse       = regexp.MustCompile(` <span class="type">朝向：</span>\s+<span class="info">([^<]+)</span>`)  // 朝向
	LoftParse         = regexp.MustCompile(`<span class="type">楼层：</span>\s+<span class="info">([^<]+)</span>`)  // 楼层
	DecorateParse     = regexp.MustCompile(`<span class="type">装修：</span>\s+<span class="info">([^<]+)</span>`)  //装修
	HouseTypeParse    = regexp.MustCompile(`<span class="type">类型：</span>\s+<span class="info">([^<]+)</span>`)  //类型： 普通住宅
	PublicTimeParse   = regexp.MustCompile(`<div class="right-info"><span id="houseCode">房屋编码：(\d+)，</span>发布时间：([^<]+)</div>`)  //房屋编号 发布时间
	PriceParse        = regexp.MustCompile(`<span class="price"><em>(\d+)</em>元/月</span>`) //租金
)

//房源详细信息
func ParserHouse(contents string)engine.RequestResult{
	var house model.House
	house.BuildingName = extendString(BuildingNameParse,contents,1)
	house.PayType = extendString(PayTypeParse,contents,1)
	house.UnitType = extendString(UnitTypeParse,contents,1)
	house.Area = extendInt(AreaParse,contents,1)
	house.Toword = extendString(TowordParse,contents,1)
	house.Loft = extendString(LoftParse,contents,1)
	house.Decorate = extendString(DecorateParse,contents,1)
	house.HouseType = extendString(HouseTypeParse,contents,1)
	house.HouseNo = extendString(PublicTimeParse,contents,1)
	house.PublicTime = extendString(PublicTimeParse,contents,2)
	house.Price = extendFloat(PriceParse,contents,1)

	var res engine.RequestResult
	res.R = append(res.R,engine.Request{})
	res.Items = house
	fmt.Printf("房源信息：%+v \n",house)
	return res
}

func extendString(r *regexp.Regexp,s string ,index int) string{
	res := r.FindStringSubmatch(s)
	if len(res) < index {
		return ""
	}
	return res[index]
}

func extendInt(r *regexp.Regexp,s string,index int) int {
	res := r.FindStringSubmatch(s)
	if len(res) < index {
		return 0
	}
	n,err := strconv.Atoi(res[index])
	if err != nil {
		return 0
	}
	return n
}

func extendFloat(r *regexp.Regexp,s string,index int)float32{
	res := r.FindStringSubmatch(s)
	if len(res) < index {
		return 0
	}
	n,err := strconv.ParseFloat(res[index],64)
	if err != nil {
		return 0
	}
	return float32(n)
}