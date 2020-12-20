package model

type ArticleItem struct {
	PageLink  string `json:"pageLink"`
	ImageLink string `json:"imageLink"`
	Topic     string `json:"topic" example:"E-health"`
	Title     string `json:"title" example:"Philips koopt Amerikaanse BioTelemetry voor 2,3 miljard euro"`
	Teaser    string `json:"teaser" example:"BioTelemetry maakt systemen om hartpatiënten te diagnosticeren en op afstand in de gaten te houden. Philips zag de vraag naar zorg op afstand toenemen vanwege de pandemie."`
}
