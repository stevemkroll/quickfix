package pet

var (
	Nameickname string
)

var Color string

type Pet struct {
	Nickname string `json:"nickname"`
	Color    string `json:"color"`
}
