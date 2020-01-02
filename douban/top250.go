package douban

import (
	"fmt"
	"go-spider/models"
	"go-spider/pkg/common"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//var wm sync.Map

func StartTop250() {
	t1 := time.Now()

	var wg sync.WaitGroup
	perpage := 25
	maxpage := 10
	for i := 0; i < maxpage; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			parse(perpage, i)
		}(i)
	}
	wg.Wait()

	// 按顺序插入
	//for i := 1; i <= maxpage*perpage; i++ {
	//	key := strconv.Itoa(i)
	//	v, ok := wm.Load(key)
	//	if !ok {
	//		continue
	//	}
	//	doubanTop250, ok := v.(models.DoubanTop250)
	//	if !ok {
	//		continue
	//	}
	//
	//	err := models.CreateDoubanTop250(&doubanTop250)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)
}

func parse(perpage, page int) {
	start := strconv.Itoa(page * perpage)
	url := "https://movie.douban.com/top250?start=" + start + "&filter="
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

	doc.Find("#content div.grid-16-8 div.article ol li").Each(func(i int, s *goquery.Selection) {
		name = s.Find("div.item .info  span.title:nth-child(1)").Text()
		star = s.Find("div.item div.info div.bd .star .rating_num").Text()
		number = s.Find("div.item div.info div.bd .star span:nth-child(4)").Text()
		desc = s.Find("div.item div.info div.bd .quote span.inq").Text()

		// 匹配 23456人评价中的人数
		numberInt, err := findNumber(number)
		if err != nil {
			log.Fatal(err)
		}

		doubanTop250 := &models.DoubanTop250{
			Sort:   page*perpage + i + 1,
			Name:   name,
			Star:   star,
			Number: numberInt,
			Desc:   desc,
			Url:    url,
		}

		//wm.Store(strconv.Itoa(doubanTop250.Sort), doubanTop250)
		if err := models.CreateDoubanTop250(doubanTop250); err != nil {
			log.Fatal(err)
		}
	})

}

// 匹配 "23456人评价" 中的人数
func findNumber(number string) (int, error) {
	// 正则
	re, err := regexp.Compile(`^\d`)
	if err != nil {
		return 0, err
	}
	substr := re.FindString(number)

	numberInt, err := strconv.Atoi(substr)
	if err != nil {
		return 0, err
	}

	return numberInt, nil
}
