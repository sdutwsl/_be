package main

import (
	"_be/config"
	"_be/service/crawler"
	"_be/service/ywq"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// method: GET
	// query: dmhy_s 需要搜索的关键词
	// response: 符合关键词的 magnet 的合集，以换行隔开
	r.GET("/GetDMHYMagnets", func(c *gin.Context) {
		var search_string = c.Request.URL.Query().Get("dmhy_s")
		c.String(200, crawler.GetDMHYMagnets(search_string))
	})

	r.GET("/GetAllRecords", func(c *gin.Context) {
		var rs = ywq.GetAllRecords()
		var json, _ = json.Marshal(rs)
		c.JSON(200, string(json))
	})

	//使用不同的端口
	if gin.Mode() == "release" {
		r.Run(config.RELEASE_CONFIG)
	} else {
		r.Run(config.DEVELOP_CONFIG)
	}
}
