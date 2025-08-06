package main

import (
	"log"
	"net"
	"regexp"
)

func checkInput(entered string) [][]string {

	ipRegex := `\b(?:\d{1,3}\.){3}\d{1,3}\b`
	re := regexp.MustCompile(ipRegex)

	matches := re.FindAllString(entered, -1)
	//for i := range matches {
	//fmt.Printf("Matches %v: %v\n", i, matches[i])
	//	}

	var ips []string
	for _, ip := range matches {
		if net.ParseIP(ip) != nil {
			ips = append(ips, string(ip))
		}
	}
	log.Println(ips)
	returnSlice := createSlice(ips)
	return returnSlice

}
