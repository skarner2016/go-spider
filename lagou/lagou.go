package lagou

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var uas = [...]string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
}

//var languages = []string{"php", "go"}
//
//var citys = []string{"北京", "上海", "广州", "深圳", "杭州", "杭州", "成都", "武汉", "江苏"}

type jobData struct {
	City     string
	Language string
}

type jobService struct {
	City     string
	Language string
}

func Start() {
	jobService := jobService{
		City:     "成都",
		Language: "php",
	}

	data, err := jobService.jobList(1)

	fmt.Println(data, err)

	//Parse()
}

// 获取 cookie
func (l *jobService) getCookie() ([]*http.Cookie, error) {
	client := http.Client{}
	formData := strings.NewReader("")
	// 1. 获取cookie
	req, err := http.NewRequest("GET", "https://www.lagou.com/gongsi/", formData)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()

	return cookies, nil
}

func (l *jobService) jobList(pn int) ([]jobData, error) {
	cookies, err := l.getCookie()
	if err != nil {
		return nil, err
	}

	// 2. 带 cookie 请求接口
	client := http.Client{}
	bs := fmt.Sprintf("first=true&pn=%d&kd=%s", pn, l.Language)
	formData := strings.NewReader(bs)
	req, err := http.NewRequest("POST", l.buildUrl(), formData)
	if err != nil {
		return nil, err
	}
	// 设置 cookie
	for _, v := range cookies {
		req.AddCookie(v)
	}
	// 设置 header
	req.Header.Add("User-Agent", l.getUserAgent())
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_%E6%95%B0%E6%8D%AE%E5%88%86%E6%9E%90?city=%E5%85%A8%E5%9B%BD&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(body)
	//var resujobData

	fmt.Println(string(body))

	//err = json.Marshal([]byte(body), &results)

	return nil, nil

}

func (l *jobService) buildUrl() string {
	prefix := "https://www.lagou.com/jobs/positionAjax.json?px=default&city="
	suffix := "&needAddtionalResult=false"

	buffer := bytes.Buffer{}
	buffer.WriteString(prefix)
	buffer.WriteString(l.City)
	buffer.WriteString(suffix)
	u, _ := url.Parse(buffer.String())
	query := u.Query()
	u.RawQuery = query.Encode()

	return u.String()
}

// 随机请求头
func (l *jobService) getUserAgent() string {
	n := rand.Intn(len(uas))

	return uas[n]
}
