package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	. "nrcnewsapi/api/nrcnewsapi/src/config"
	. "nrcnewsapi/api/nrcnewsapi/src/model"
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

			linkOfPage, _ := goQuerySelection.
				Find("a").
				Attr("href")

			imageLink := strings.
				Split(e.ChildAttr("img", "data-src"), "|")[0]

			header := goQuerySelection.Find(".nmt-item__content")

			topic := trimText(header.Find("h6").Text())
			title := trimText(header.Find("h3").Text())
			teaser := trimText(header.Find(".nmt-item__teaser").Text())

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

		log.Println("Article items scraped:", len(articleList))

		context.JSON(http.StatusOK, articleList)
	}
}

func (scraper Scraper) GetArticle() gin.HandlerFunc {
	return func(context *gin.Context) {
		var articleItem ArticleItem
		context.BindJSON(&articleItem)

		c := GetCollector()

		initializeCalls(c)

		var article Article

		var sectionList []Section

		var buff []Dummy

		c.OnHTML("div.content", func(e *colly.HTMLElement) {

			goQuerySelection := e.DOM

			var dummy Dummy
			goQuerySelection.Find("div.content > p, div.content > h2").
				Each(func(i int, selection *goquery.Selection) {

					if selection.Parent().Is("aside") {
						return
					}

					if selection.Is("h2") {
						dummy.Title = selection.Text()
						dummy.Content = ""
						dummy.Type = "h2"

						buff = append(buff, dummy)

					} else if selection.Is("p") && len(selection.Text()) > 0 {
						dummy.Content = selection.Text()
						dummy.Type = "p"

						buff = append(buff, dummy)
					}
				})

			groupedByTitle := linq.From(buff).GroupBy(
				func(i interface{}) interface{} { return i.(Dummy).Title },
				func(i interface{}) interface{} { return i.(Dummy).Content })

			groupedByTitle.ForEach(func(i interface{}) {
				var section Section
				section.Title = i.(linq.Group).
					Key.(string)

				linq.From(i.(linq.Group).Group).
					ToSlice(&section.Contents)

				sectionList = append(sectionList, section)
			})

			article = Article{
				ArticleItem: ArticleItem{
					PageLink:  articleItem.PageLink,
					ImageLink: articleItem.ImageLink,
					Topic:     articleItem.Topic,
					Title:     articleItem.Title,
					Teaser:    articleItem.Teaser,
				},
				SectionList: sectionList,
			}
		})

		c.Visit(articleItem.PageLink)

		c.Wait()

		log.Println("Article scraped succesfully")

		context.JSON(http.StatusOK, article)

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

func trimText(text string) string {
	trimmedSpaces := strings.TrimSpace(text)
	trimmedLeft := strings.TrimLeft(trimmedSpaces, " ")
	trimmedFinal := strings.TrimRight(trimmedLeft, " ")
	return trimmedFinal
}

func printSection(section Section) {
	log.Println("Section title: {}", section.Title)

	for _, content := range section.Contents {
		log.Println("Section content: {}", content)
	}

	log.Println("-------------------")
}

type Dummy struct {
	Title   string
	Content string
	Type    string
}
