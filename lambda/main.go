package main

import (
	"fmt"
	"log"
	"net/url"

	vt "github.com/VirusTotal/vt-go"
	"github.com/likexian/whois-go"
)

func main() {
	result, err := whois.Whois("likexian.com")
	if err == nil {
		fmt.Println(result)
	}
	var apikey = "a749ffc43346e9086a589f9758b0312008b8bcbd62fb9ee85a7fc6faa67c565d"
	client := vt.NewClient(apikey)
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
