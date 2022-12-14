package crawler

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//查询字符串名和动漫花园域名
const DMHY_SEARCH_DOMAIN = "https://dmhy.anoneko.com/topics/list"

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
