package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type currentjson struct {
	Id   string  `json:"server"`
	Content string `json:"coder"`
}

func handler_json(writer http.ResponseWriter, request *http.Request) {

	db, _ := gorm.Open("mysql", "root:root@/test-api?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	defer db.Close()

	var tmp currentjson
	server := request.URL.Query()["server"];
	db.First(&tmp, "Id = ?", server[0])

	writer.Write([]byte(tmp.Content))

}

func handler(writer http.ResponseWriter, request *http.Request) {

	db, _ := gorm.Open("mysql", "root:root@/test-api?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	defer db.Close()

	code := request.URL.Query()["code"];
	server := request.URL.Query()["server"];

	var tmp currentjson

	// check if id exists
	if db.Where("Id = ?", server[0]).First(&tmp).RecordNotFound() {
		//or create a new ID
		guid := xid.New()
		db.Create(&currentjson{Content: code[0], Id: guid.String()})
		writer.Write([]byte(tmp.Content))

	}

	// automatic change the id
	if err := db.Model(&tmp).Where("Id = ?", server[0]).Update("Content", code[0]).Error; err != nil{
	}

}

func main() {

	http.HandleFunc("/json", handler_json)
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8077", nil)
	}
