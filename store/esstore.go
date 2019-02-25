package store

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"house_spider/house/model"
	"log"
	"os"
	"reflect"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200"

/**
  初始化链接
 */
func init() {
    errorlog := log.New(os.Stderr,"app",log.LstdFlags)
    var err error
    client,err = elastic.NewClient(elastic.SetErrorLog(errorlog),elastic.SetURL(host))
    if err != nil {
    	panic(err)
	}
}

/**
  BuildingName string //小区名称
	PayType string  // 付3押1
	UnitType string // 3室2厅2卫
	Area int //面积
	Toword string // 朝向
	Loft string // 楼层
	Decorate string //装修
	HouseType string //类型： 普通住宅
	PublicTime string //发布时间
	Price float32 //租金
	HouseNo string //房屋编号
 */

func Save(item interface{}) error {
	 newItem ,ok := item.(model.House)
	 if !ok {
	 	return nil
	 }
	 res ,err := client.Index().Index("zufang").Type("house").Id(newItem.HouseNo).BodyJson(newItem).Do(context.Background())
	 if err != nil {
	 	log.Printf("save err ,error is %v\n",err)
	 	return nil
	 }
	_, err = client.Flush().Index("zufang").Do(context.TODO())
	if err != nil {
		log.Printf("save err ,error is %v\n",err)
		return nil
	}
	 log.Printf("save success id is %s info is %v \n",res.Id,item)
	 return  nil
}

func Get(keyword string) []model.House {

	var res *elastic.SearchResult
	var err error
	// Specify highlighter
	hl := elastic.NewHighlight()
	hl = hl.Fields(elastic.NewHighlighterField("BuildingName"))
	hl = hl.PreTags(`<em style="color:red">`).PostTags("</em>")

	// Match all should return all documents
	query := elastic.NewPrefixQuery("BuildingName", keyword)
	res, err = client.Search().
		Index("zufang").
		Highlight(hl).
		Query(query).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		log.Printf("query err ,error is %v\n",err)
		return nil
	}
	if res.Hits == nil {
		log.Printf("query nil ")
		return nil
	}
	var typ model.House
	var data []model.House
	mapdata := make(map[string]string)
	// 查询到高亮部分
	for _,v := range res.Hits.Hits{
		if err := json.Unmarshal(*v.Source, &typ); err != nil {
			log.Printf("error is %v\n",err)
			continue
		}
		if v.Highlight == nil || len(v.Highlight) == 0 {
			log.Printf("error is %v\n",err)
			continue
		}
		mapdata[v.Id] = v.Highlight["BuildingName"][0]
	}
	//追加到结果里
	for _,item := range res.Each(reflect.TypeOf(typ)){
		log.Printf("item %v\n",item)
		t := item.(model.House)
		if name,ok := mapdata[t.HouseNo] ; ok {
			t.BuildingName = name
		}
		data = append(data,t)
	}

	return data
}
