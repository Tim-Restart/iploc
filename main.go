package main

import (
	"fmt"
	"os"
)





func main() {


	
	queries := len(os.Args)
	switch {
	case queries == 1:
		fmt.Println("no IP address provided")
		os.Exit(1)
	case queries == 2:
		ip := os.Args[1]
		result, err := getIP(ip)
		if err != nil {
			fmt.Println("Error sending request")
			fmt.Println(err)
			return
		}
		fmt.Println(result.Zip)
		fmt.Println(result.City)
		fmt.Println(result.Region)
	case queries > 2:
		i := 1
		for i < len(os.Args) {
			ip := os.Args[i]
			result, err := getIP(ip)
			if err != nil {
				fmt.Println("Error sending request")
				fmt.Println(err)
				return
			}
			fmt.Println(result.Zip)
			fmt.Println(result.City)
			fmt.Println(result.Region)
			i++
		}
	default:
		fmt.Println("Incorrect queries, please try again")
	}
	}

