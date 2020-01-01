package douban

import (
	"fmt"
	"go-spider/models"
	"go-spider/pkg/common"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func StartTop250() {
	t1 := time.Now() // get current time
	// TODO: 查询所有 top250
	parse()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)
}

func parse() {
	url := "https://movie.douban.com/top250?start=" + "50" + "&filter="
	response, err := common.SendRequest(map[string]string{}, url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("status code err: %d %s", response.StatusCode, response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var name, star, number, desc string

	// TODO: 优化(批量插入)
	doc.Find("#content div.grid-16-8 div.article ol li").Each(func(i int, s *goquery.Selection) {
		name = s.Find("div.item .info  span.title:nth-child(1)").Text()
		star = s.Find("div.item div.info div.bd .star .rating_num").Text()
		number = s.Find("div.item div.info div.bd .star span:nth-child(4)").Text()
		desc = s.Find("div.item div.info div.bd .quote span.inq").Text()

		doubanTop250 := &models.DoubanTop250{
			Name:   name,
			Star:   star,
			Number: number,
			Desc:   desc,
			Url:    url,
		}
		if err := models.CreateDoubanTop250(doubanTop250); err != nil {
			log.Fatal(err)
		}
	})
}
