package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	wordPtr *string
)

func getCrtSh(host string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://crt.sh/?output=json", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query() // Get a copy of the query values.
	q.Add("q", host)     // Add a new value to the set.
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

	fmt.Println(q.Encode(), err)
}

func init() {
	wordPtr = flag.String("h", "facebook.com", "hostname")
}

func main() {
	flag.Parse()
	fmt.Println(*wordPtr)
	getCrtSh(*wordPtr)
}

//https://crt.sh/?q=%.facebook.com&output=json
