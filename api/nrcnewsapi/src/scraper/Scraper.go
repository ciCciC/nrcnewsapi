package scraper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	. "nrcnewsapi/api/nrcnewsapi/src/config"
	. "nrcnewsapi/api/nrcnewsapi/src/model"
	"log"
	"net/http"
	"strings"
)

type Scraper struct {
	Endpoint string
}

func (scraper Scraper) GetAllArticles() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := GetCollector()

		initializeCalls(c)

		var articleList []ArticleItem

		c.OnHTML("div.nmt-item__inner", func(e *colly.HTMLElement) {

			goQuerySelection := e.DOM

			linkOfPage, _ := goQuerySelection.Find("a").Attr("href")
			imageLink := strings.Split(e.ChildAttr("img", "data-src"), "|")[0]

			header := goQuerySelection.Find(".nmt-item__content")

			topic := strings.TrimSpace(header.Find("h6").Text())
			title := header.Find("h3").Text()
			teaser := header.Find(".nmt-item__teaser").Text()

			articleList = append(articleList,
				ArticleItem{
					PageLink:  API + linkOfPage,
					ImageLink: imageLink,
					Topic:     topic,
					Title:     title,
					Teaser:    teaser})
		})

		c.Visit(API + "/categorie/" + scraper.Endpoint)

		c.Wait()

		log.Println("Articles scraped:", len(articleList))

		context.JSON(http.StatusOK, articleList)
	}
}

func initializeCalls(c *colly.Collector) {
	// Before making a request ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
}