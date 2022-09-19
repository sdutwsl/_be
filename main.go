package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const QUERY_NAME = "dmhy_s"
const DMHY_SEARCH_DOMAIN = "https://dmhy.anoneko.com/topics/list"

func main() {
	r := gin.Default()
	r.GET("/GetDMHYMagnets", func(c *gin.Context) {
		var search_string = c.Request.URL.Query().Get(QUERY_NAME)
		c.String(200, GetDMHYMagnets(search_string))
	})

	r.Run(":7003")
}

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
