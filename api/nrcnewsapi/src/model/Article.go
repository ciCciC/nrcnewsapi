package model

type Article struct {
	ArticleItem
	SectionList []Section `json:"sectionList"`
}
