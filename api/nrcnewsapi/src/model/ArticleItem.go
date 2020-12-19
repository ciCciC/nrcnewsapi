package model

type ArticleItem struct {
	PageLink  string `json:"pageLink" example:"page link"`
	ImageLink string `json:"imageLink" example:"page link"`
	Topic     string `json:"topic" example:"page link"`
	Title     string `json:"title" example:"page link"`
	Teaser    string `json:"teaser" example:"page link"`
}
