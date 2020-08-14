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
	"nrcnewsapi/api/nrcnewsapi/src/util"
	"strings"
)

type Scraper struct {
	Endpoint  string
	Topic     string
	Endpoints []string
}

func (scraper Scraper) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := GetCollector()

		cacheKey := strings.Join(scraper.Endpoints, "")

		var articleList []ArticleItem

		cachedArticles, found := util.GetCache(cacheKey)

		if found {

			log.Println("Fetches cached articles")
			context.JSON(http.StatusOK, cachedArticles)

		} else {

			initializeCalls(c)

			for _, endpoint := range scraper.Endpoints {

				c.OnHTML("div.nmt-item__inner", func(e *colly.HTMLElement) {

					goQuerySelection := e.DOM

					linkOfPage, _ := goQuerySelection.
						Find("a").
						Attr("href")

					imageLink := strings.
						Split(e.ChildAttr(IMG, "data-src"), "|")[0]

					header := goQuerySelection.Find(".nmt-item__content")

					topic := util.TrimText(header.Find("h6").Text())
					title := util.TrimText(header.Find("h3").Text())
					teaser := util.TrimText(header.Find(".nmt-item__teaser").Text())

					articleList = append(articleList,
						ArticleItem{
							PageLink:  API + linkOfPage,
							ImageLink: imageLink,
							Topic:     topic,
							Title:     title,
							Teaser:    teaser})
				})

				c.Visit(API + "/categorie/" + endpoint)
				c.Wait()

				log.Println("Article items scraped:", len(articleList))

			}

			util.SetCache(cacheKey, articleList)

			context.JSON(http.StatusOK, articleList)
		}
	}
}

func (scraper Scraper) GetAllArticles() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := GetCollector()

		var articleList []ArticleItem

		cachedArticles, found := util.GetCache(scraper.Endpoint)

		if found {

			log.Println("Fetches cached articles")
			context.JSON(http.StatusOK, cachedArticles)

		} else {

			initializeCalls(c)

			c.OnHTML("div.nmt-item__inner", func(e *colly.HTMLElement) {

				goQuerySelection := e.DOM

				linkOfPage, _ := goQuerySelection.
					Find("a").
					Attr("href")

				imageLink := strings.
					Split(e.ChildAttr(IMG, "data-src"), "|")[0]

				header := goQuerySelection.Find(".nmt-item__content")

				topic := util.TrimText(header.Find("h6").Text())
				title := util.TrimText(header.Find("h3").Text())
				teaser := util.TrimText(header.Find(".nmt-item__teaser").Text())

				if util.IsEmpty(topic) {
					topic = scraper.Topic
				}

				articleList = append(articleList,
					ArticleItem{
						PageLink:  API + linkOfPage,
						ImageLink: imageLink,
						Topic:     topic,
						Title:     title,
						Teaser:    teaser})
			})

			c.Visit(API + scraper.Endpoint)

			c.Wait()

			log.Println("Article items scraped:", len(articleList))

			util.SetCache(scraper.Endpoint, articleList)

			context.JSON(http.StatusOK, articleList)
		}
	}
}

func (scraper Scraper) GetArticle() gin.HandlerFunc {
	return func(context *gin.Context) {
		var articleItem ArticleItem
		context.BindJSON(&articleItem)

		var article Article
		var sectionList []Section
		var buff []Dummy

		c := GetCollector()

		initializeCalls(c)

		c.OnHTML("div.content", func(e *colly.HTMLElement) {

			goQuerySelection := e.DOM

			var dummy Dummy
			goQuerySelection.Find("div.content > p, div.content > h2, div.content > figure").
				Each(func(i int, selection *goquery.Selection) {

					if selection.Parent().Is("aside") {
						return
					}

					if selection.Is(H2) {
						dummy.Title = selection.Text()
						dummy.ContentBody.Content = ""
						dummy.ContentBody.CType = H2

					} else if selection.Is(P) && len(selection.Text()) > 0 {
						dummy.ContentBody.Content = selection.Text()
						dummy.ContentBody.CType = P

					} else if selection.Is(FIGURE) {
						attr := e.ChildAttr(IMG, "data-src")
						dummy.ContentBody.Content = strings.Split(attr, "|")[0]
						dummy.ContentBody.CType = IMG
					}

					buff = append(buff, dummy)
				})

			groupedByTitle := linq.From(buff).GroupBy(
				func(i interface{}) interface{} { return i.(Dummy).Title },
				func(i interface{}) interface{} { return i.(Dummy).ContentBody })

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

		log.Println("Fetched Article: ", article.Topic, article.Title)
		log.Println("Article scraped successfully")

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

func printSection(section Section) {
	log.Println("Section title: {}", section.Title)

	for _, content := range section.Contents {
		log.Println("Section content: {}", content)
	}

	log.Println("-------------------")
}

type Dummy struct {
	Title       string
	ContentBody ContentBody
}

const (
	H2     = "h2"
	P      = "p"
	IMG    = "img"
	FIGURE = "figure"
)
