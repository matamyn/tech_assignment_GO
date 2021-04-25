package model

var DefaultLink = "https.short_link.ru/"

type LinkDbRow struct {
	ID_           uint
	ShortLinkKey_ string
	Link_         string
}
