# Welcome to iploc

iploc is a command line tool used to identify IP address locations. It can take a single IP address or multiple IP addresses.

Usage is currently limited to 45 requests per minute using ip-api.com's free calls.

## Installation
Currently this is only supported on linux, or windows through WSL.

A copy of the binary is avaliable in the repo called iploc.

Only standard packages have been used for this program currently.

To download a local copy and build it, just run `gh repo clone Tim-Restart/iploc` if you have the github CLI, or other methods avaliable on the green code button.

Then run: 
`go build .`
From the root directory for the project.

Usage from the directory will be:
`./iploc [ip address]`

To make it able to be run from any folder location:
`sudo mv iploc /usr/local/bin`

This will allow you to run it as:
`iploc [ip address]` 
from any folder in the CLI.

## Single IP
Usage: `iploc [ip address]`
This will conduct a query on a single IP address and return the results

## Multiple IP
Usage: `iploc [ip address] [ip address]`
This will return multiple results for entered IP addresses

When entering multiple IP addresses, commas can be at the end, eg: 100.100.01.1,
This will still create the correct query.

If a result is empty, this means there was not data for that field from the API call.

Data can be saved to a file using `> report.txt` or similar