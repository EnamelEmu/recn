package subdomain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"log"
	"bufio"
	"net/http"
	"os"
	//"encoding/json"
)

type CrtJson struct {
	IssuerCaID        int    `json:"issuer_ca_id"`
	IssuerName        string `json:"issuer_name"`
	NameValue         string `json:"name_value"`
	MinCertID         int    `json:"min_cert_id"`
	MinEntryTimestamp string `json:"min_entry_timestamp"`
	NotBefore         string `json:"not_before"`
	NotAfter          string `json:"not_after"`
}

type Crts []*CrtJson

func (crt CrtJson) showFullCTlog() {
	fmt.Printf("%s\n", crt.NameValue) //dns endpoint
}

func getCrtSh(host string) ([]byte, error) {
	resp, err := http.Get("https://crt.sh/?q=" + host + "&output=json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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

func parseCrtQr(query string) error {
	var crts Crts

	resp, err := getCrtSh(query)
	if err = json.Unmarshal(resp, &crts); err != nil {
		return err
	}
	for _, crt := range crts {
		crt.showFullCTlog()
	}

	return nil
}
func main() {
	/*wordPtr := flag.String("h", "facebook.com", "hostname")
	flag.Parse()
	fmt.Println(*wordPtr) */
	stat, _ := os.Stdin.Stat()
	if(stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No stdin input detected")
		os.Exit(1)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		url := s.Text()
		parseCrtQr(url)
	}

}

//https://crt.sh/?q=%.facebook.com&output=json
