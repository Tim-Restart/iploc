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

func resultsToSlice(res *IPResults, resData map[string]string) map[string]string {
	// take the IP results
	// add them to the map
	// return the map

	resData["1. IP Address"] = res.Query
	resData["2. Country"] = res.Country
	resData["3. Region"] = res.Region
	resData["4. City"] = res.City
	resData["5. ISP"] = res.Isp

	if res.Mobile {
		resData["6. Mobile"] = "True"
	} else {
		resData["6. Mobile"] = "False"
	}
	if res.Proxy {
		resData["7. Proxy/VPN"] = "True"
	} else {
		resData["7. Proxy/VPN"] = "False"
	}

	return resData
}
