package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"_be/be_config"

	"github.com/gin-gonic/gin"
)

//查询字符串名和动漫花园域名
const QUERY_NAME = "dmhy_s"
const DMHY_SEARCH_DOMAIN = "https://dmhy.anoneko.com/topics/list"
const RELEASE_CONFIG = ":19971"
const DEVELOP_CONFIG = ":7003"

func main() {
	fmt.Println(be_config.MYSQL_HOST)
	r := gin.Default()
	r.GET("/GetDMHYMagnets", func(c *gin.Context) {
		var search_string = c.Request.URL.Query().Get(QUERY_NAME)
		c.String(200, GetDMHYMagnets(search_string))
	})
	//使用不同的端口
	if gin.Mode() == "release" {
		r.Run(RELEASE_CONFIG)
	} else {
		r.Run(DEVELOP_CONFIG)
	}
}

//获取动漫花园magnet链接
func GetDMHYMagnets(s string) string {
	var req, err = http.NewRequest("GET", DMHY_SEARCH_DOMAIN, nil)
	if err != nil {
		return "request error"
	}
	var q = req.URL.Query()
	q.Add("keyword", s)
	req.URL.RawQuery = q.Encode()
	var res, _ = http.DefaultClient.Do(req)
	var body, _ = ioutil.ReadAll(res.Body)
	var re = regexp.MustCompile("\"magnet.*\"")
	var magnets = re.FindAllString(string(body), -1)
	for i, v := range magnets {
		magnets[i] = v[1 : len(v)-1]
	}
	return strings.Join(magnets, "\n")
}
