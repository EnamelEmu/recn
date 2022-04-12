package dork

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	cseapi string
	cxapi string
)

type gglConfig struct {
	cseapiConf string
	cxapiConf string
}

func helper() {
	fmt.Println("Usage: dork -a [google CSE api key] -c [config file] [string]")
}


func readapiandcx(input gglConfig) {
	cseapi = input.cseapiConf
	cxapi = input.cxapiConf
}

// exact terms, exclude terms, filetype, linksite, sitesearch

func getCse(query string) () {
	resp, err := http.Get("https://customsearch.googleapis.com/customsearch/v1?key=" + apikey + "&cx=" + cxid + "&q=" + query)
}
