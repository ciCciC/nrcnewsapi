package model

type ArticleItem struct {
	PageLink  string `json:"pageLink"`
	ImageLink string `json:"imageLink"`
	Topic     string `json:"topic"`
	Title     string `json:"title"`
	Teaser    string `json:"teaser"`
}
