package main

import (
	"fmt"
)

func report(result *IPResults) error {
	fmt.Printf(`
######################################
Results
IP: %v
City: %v 
Region: %v 
Lat/Long: %v, %v
ISP: %v 
Hosting: %v 
Org: %v 
Proxy: %v 
Mobile: %v 
Reverse DNS: %v
######################################
`, result.Query, result.City, result.Region, result.Lat, result.Lon, result.Isp, result.Hosting, result.Org, result.Proxy, result.Mobile, result.Reverse)
	return nil
}


func help() {
			fmt.Printf(`
#################
Welcome to ipLoc:
#################
Single IP
Usage: iploc [ip address]
This will conduct a query on a single IP address and return the results

Multiple IP
Usage: iploc [ip address] [ip address]
This will return multiple results for entered IP addresses

When entering multiple IP addresses, commas can be at the end, eg: 100.100.01.1,
This will still create the correct query.

Usage is currently limited to 45 requests per minute using ip-api.com's free calls.

If a result is empty, this means there was not data for that field from the API call.

Data can be saved to a file using > report.txt or similar
`)
	return
}