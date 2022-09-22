package controller

import (
	"_be/config"
	"_be/service/crawler"
	"_be/service/ywq"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func RunApp() {

	r := gin.Default()

	// method: GET
	// query: dmhy_s 需要搜索的关键词
	// response: 符合关键词的 magnet 的合集，以换行隔开
	// 搜索dmhy所有符合条件的种子连接并返回
	r.GET("/GetDMHYMagnets", func(c *gin.Context) {
		var search_string = c.Request.URL.Query().Get("dmhy_s")
		c.String(200, crawler.GetDMHYMagnets(search_string))
	})

	// method: GET
	// response: 所有的时间记录
	// 获取所有的时间记录
	r.GET("/GetAllRecords", func(c *gin.Context) {
		var rs = ywq.GetAllRecords()
		var json, _ = json.Marshal(rs)
		c.JSON(200, string(json))
	})

	// method: GET
	// response: 添加时间的记录的id
	// 增加一条时间的id
	r.GET("/AddRecord", func(c *gin.Context) {
		var id = ywq.AddRecord()
		c.JSON(200, id)
	})

	r.Run(config.GIN_CONFIG)
}
