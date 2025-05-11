package airport

type AirportCodeSearchResult struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type AirportCodeSearchResponse struct {
	Results []AirportCodeSearchResult `json:"results"`
}
