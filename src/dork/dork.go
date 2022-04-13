package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	cseapi string
	cxid string
)

type gglConfig struct {
	CseapiConf string `json:"cseapi"`
	Cxidconf string `json:"cxid"`
}

type gglResponseItems struct {
	Links []struct {
		Link string `json:"link"`
		Title string `json:"title"`
		json.RawMessage
	} `json:"items"`
}

type Links []*gglResponseItems

func (link gglResponseItems) retValues() string {
	return link.Links[0].Link + link.Links[0].Title
}

func helper() {
	fmt.Println("Usage: dork -c [config file] [string]")
}

func readapiandcx(path string) error {
	if path == "" {
		helper()
		return nil
	}
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Erraor:", err)
	}
	inputJson := []byte(dat)
	var xgglconfig gglConfig
	err2 := json.Unmarshal(inputJson, &xgglconfig)
	if err2 != nil {
		fmt.Println(string(inputJson))
		fmt.Println("Error: ", err2)
	}
	cseapi = xgglconfig.CseapiConf
	cxid = xgglconfig.Cxidconf
	return nil
}

// exact terms, exclude terms, filetype, linksite, sitesearch

func getCse(query string) ([]byte, error) {

	resp, err := http.Get("https://customsearch.googleapis.com/customsearch/v1?key=" + cseapi + "&cx=" + cxid + "&filter=0" + "&exactTerms=" + query)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return body, nil
}

func parseCse(query string) error {
	var linktr gglResponseItems

	resp, err := getCse(query)
	if err = json.Unmarshal(resp, &linktr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return  err
	}
	for _, linkvalue := range linktr.Links {
		fmt.Println(linkvalue.Link, linkvalue.Title)
	}

	return nil
}

func main() {
	var cFlag = flag.String("c", "", "config file")
	flag.Parse()
	readapiandcx(*cFlag)
	lastarg := os.Args[len(os.Args)-1]
	parseCse(lastarg)
}
