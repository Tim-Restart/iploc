package main

import (
	"fmt"
	"os"
)





func main() {



	switch len(os.Args) {
	case 1:
		fmt.Println("no IP address provided")
		os.Exit(1)
	case 2:
		ip := os.Args[1]
		result, err := getIP(ip)
		if err != nil {
			fmt.Println("Error sending request")
			fmt.Println(err)
			return
		}
		fmt.Println(result.City)
		
		}
	}

