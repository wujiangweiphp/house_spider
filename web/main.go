package main

import (
	"encoding/json"
	"house_spider/house/model"
	"house_spider/store"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
)

func search(writer http.ResponseWriter, request *http.Request) {
	var data []model.House
	var err error
	err = request.ParseForm()
	if keyword,ok := request.Form["keyword"]; ok && err == nil {
		data = store.Get(keyword[0])
	}
	jsdata,_:= json.Marshal(data)
	writer.Header().Add("Content-type","application/json")
	writer.Write(jsdata)
}

func house(writer http.ResponseWriter, request *http.Request)  {
	/*_, path, _, ok := runtime.Caller(1)
	if !ok {
		log.Printf("file error ")
	}*/
	path := `E:\goproject\src\house_spider\web`
	t := template.New("index.html")
	t = template.Must(t.ParseFiles(path + "/view/index.html"))
	data := struct {
		Title string
	}{
		"杭州租房信息查询",
	}
	t.Execute(writer,data)
}

type handleRequest func(writer http.ResponseWriter, request *http.Request)

func warpHandle(handle handleRequest) func(http.ResponseWriter, *http.Request)  {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover();err != nil {
				log.Printf("访问出错，错误：%v\n",err)
				debug.PrintStack()
			}
		}()
		handle(writer,request)
	}
}

func main() {
	//处理搜索
	http.HandleFunc("/search", warpHandle(search))
	http.HandleFunc("/", warpHandle(house))
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Printf("server start error ,error is %v \n",err)
	}
}
