package main

import (
	"fmt"
	"net/http"
	"io"
)


// Caller function

func getIP(ip string) (*IPResults, error) {

	var ipResults IPResults
	var body []byte

	
	baseURL := "http://ip-api.com/json/"
	queries := "?fields=status,message,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,reverse,mobile,proxy,hosting,query"
	request := baseURL + ip + queries

	res, err := http.Get(request)
	if err != nil {
		fmt.Println("Error getting data")
		return nil, err
	}


	body, err = io.ReadAll(res.Body)
	if res.StatusCode <200 || res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	err = unmarshalJSON(body, &ipResults)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return &ipResults, nil

}
