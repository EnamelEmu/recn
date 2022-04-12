package dork

import (
	"bufio"
	"encoding/json"
	"flag"
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

func readapiandcx(inputJson []byte) {
	var xgglconfig gglConfig
	err := json.Unmarshal(inputJson, xgglconfig)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	cseapi = xgglconfig.cseapiConf
	cxapi = xgglconfig.cxapiConf
}

// exact terms, exclude terms, filetype, linksite, sitesearch

func getCse(query string) () {
	resp, err := http.Get("https://customsearch.googleapis.com/customsearch/v1?key=" + apikey + "&cx=" + cxid + "&q=" + query)
}

func main() {
	var cFlag = flag.String("c", "./creds.json", "cx id and api keys for google custom search engine service")
	flag.Parse()
	stat, _ := os.Stdin.Stat()
	if(stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No stdin input detected")
		os.Exit(1)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		 := s.Text()

	}
}
