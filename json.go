package main

import (
	"encoding/json"
)

type IPResults struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Reverse     string  `json:"reverse"`
	Mobile      bool    `json:"mobile"`
	Proxy       bool    `json:"proxy"`
	Hosting     bool    `json:"hosting"`
}

func unmarshalJSON(data []byte, res *IPResults) error {
	return json.Unmarshal(data, res)
}

func resultsSlice(res *IPResults, results []string) []string {

	var mobile string
	var proxy string

	if res.Mobile {
		mobile = "True"
	} else {
		mobile = "False"
	}
	if res.Proxy {
		proxy = "True"
	} else {
		proxy = "False"
	}

	// Fill slice of slices with multiple resultsSlice

	results[0] = res.Query
	results[1] = res.City
	results[2] = res.Region
	results[3] = res.Country
	results[4] = res.Isp
	results[5] = mobile
	results[6] = proxy

	return results
}
