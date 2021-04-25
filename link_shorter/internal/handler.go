package link_shorter

type Handler interface {
	GetLink() string
	CreateShortLink(string2 string) string
}
