package config

import (
	"github.com/gocolly/colly"
	. "nrcnewsapi/api/nrcnewsapi/src/util"
)

var API = "http://www.nrc.nl"

func GetCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(
			"nrc.nl",
			"http://nrc.nl",
			"https://nrc.nl",
			"www.nrc.nl"),
		colly.UserAgent(GenChromeUA()),
		colly.Async(true),
	)
	return collector
}