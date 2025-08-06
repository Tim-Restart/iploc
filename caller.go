package main

import (
	"fmt"
	"io"
	"net/http"
)

// Caller function

func getIP(ip string) (*IPResults, error) {
	//log.Print("Entered Check")
	//var ipResults IPResults
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
	if res.StatusCode < 200 || res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	r := &IPResults{}

	err = unmarshalJSON(body, r)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return r, nil

}

//func reflectResults(*IPResults) (*)
//
