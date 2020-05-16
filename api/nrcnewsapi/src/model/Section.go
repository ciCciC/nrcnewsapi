package model

type Section struct {
	Title string `json:"title"`
	//Contents []string
	Contents []ContentBody `json:"contentBody"`
}
