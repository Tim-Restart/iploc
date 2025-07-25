package main

import (
	"fmt"
	"os"
	"strings"
)


// Trim any comma's located after arguments

func commaTrim(entered string) string {
	s := strings.TrimSuffix(entered, ",")
	return s
}



func main() {


	
	queries := len(os.Args)
	switch {
	case queries == 1:
		// Calls the help file with usage from reports.go
		help()
		return
	case queries == 2:
		ip := commaTrim(os.Args[1])
		result, err := getIP(ip)
		if err != nil {
			fmt.Println("Error sending request")
			fmt.Println(err)
			return
		}
		err = report(result)
		if err  != nil {
			fmt.Println("error writing report")
		}
	case queries > 2:
		i := 1
		for i < len(os.Args) {
			ip := commaTrim(os.Args[i])
			result, err := getIP(ip)
			if err != nil {
				fmt.Println("Error sending request")
				fmt.Println(err)
				return
			}
			err = report(result)
			if err  != nil {
				fmt.Println("error writing report")
			}
			i++
		}
	default:
		fmt.Println("Incorrect queries, please try again")
	}
}

