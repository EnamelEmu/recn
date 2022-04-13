package dork

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
	cseapiConf string
	cxidConf string
}

type gglResponse struct {
	link string
	title string
}

type links []*gglResponse

func (link gglResponse) retValues() string {
	return link.link + link.title
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
		fmt.Println("Error:", err)
	}
	inputJson := []byte(dat)
	var xgglconfig gglConfig
	err2 := json.Unmarshal(inputJson, xgglconfig)
	if err2 != nil {
		fmt.Println("Error: ", err)
	}
	cseapi = xgglconfig.cseapiConf
	cxid = xgglconfig.cxidConf
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
	var linktr links
	resp, err := getCse(query)
	if err = json.Unmarshal(resp, &linktr); err != nil {
		return  err
	}
	for _, linkvalue := range linktr {
		fmt.Println(linkvalue.retValues())
	}
	return nil
}

func main() {
	var cFlag = flag.String("c", "", "config file")
	flag.Parse()

}
