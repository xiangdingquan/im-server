package api

type Question struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

type ResolveRequest struct {
	Name          string `form:"name"`
	Type          string `form:"type"`
	RandomPadding string `json:"random_padding"`
}

type ResolveResponse struct {
	Status   int        `json:"Status"`
	TC       bool       `json:"TC"`
	RD       bool       `json:"RD"`
	RA       bool       `json:"RA"`
	AD       bool       `json:"AD"`
	CD       bool       `json:"CD"`
	Question []Question `json:"Question"`
	Answer   []Answer   `json:"Answer"`
	Comment  string     `json:"comment"`
}

const (
	DnsTypeUnknown = iota
	DnsTypeA       = 1
	DnsTypeTXT     = 16
)

func GetDnsType(t2 string) int {
	switch t2 {
	case "A":
		return DnsTypeA
	case "TXT":
		return DnsTypeTXT
	default:
		return DnsTypeUnknown
	}
}
