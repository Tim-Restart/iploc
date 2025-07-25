# Welcome to ipLoc:

## Single IP
Usage: `iploc [ip address]`
This will conduct a query on a single IP address and return the results

## Multiple IP
Usage: `iploc [ip address] [ip address]`
This will return multiple results for entered IP addresses

When entering multiple IP addresses, commas can be at the end, eg: 100.100.01.1,
This will still create the correct query.

Usage is currently limited to 45 requests per minute using ip-api.com's free calls.

If a result is empty, this means there was not data for that field from the API call.

Data can be saved to a file using `> report.txt` or similar