package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "赵辉的博客"
	indexData.Desc = "博客描述"
	jsonString, _ := json.Marshal(indexData)
	w.Write(jsonString)
}
func main() {
	// 程序入口 ，宇哥项目只能有一个入口
	// web程序，http程序 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
