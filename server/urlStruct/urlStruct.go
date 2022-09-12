package urlStruct

type UrlInformation struct {
	Url            string  `json:"url,omitempty"`
	Views          float64 `json:"views,omitempty"`
	RelevanceScore float64 `json:"relevanceScore,omitempty"`
}

type UrlList struct {
	Data []UrlInformation `json:"data,omitempty"`
	Count int `json:"count,omitempty"`
}