package types

type XmlPlayer struct {
	Id        uint64 `xml:"id,attr"`
	LastLogin string `xml:"lastlogin,attr"`
}

type XmlPlayers struct {
	Players []XmlPlayer `xml:"player"`
}

type XmlItemTypes struct {
	Items []*ItemType `xml:"item"`
}

type steamPlayerBody struct {
	Response steamPlayerResponse
}

type steamPlayerResponse struct {
	Players []*steamPlayer
}

type steamPlayer struct {
	PersonaName string
	ProfileUrl  string
	AvatarUrl   string `json:"avatar"`
}
