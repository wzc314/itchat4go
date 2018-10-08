package components

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetContact() (map[string]User, error) {
	contacts := make(map[string]User)

	req, _ := http.NewRequest("GET", loginData.info["url"]+"/webwxgetcontact", nil)
	req.URL.RawQuery = url.Values{
		"lang":        {"zh_CN"},
		"pass_ticket": {loginData.info["pass_ticket"]},
		"r":           {GetTimestamp()},
		"seq":         {"0"},
		"skey":        {loginData.baseReq.Skey},
	}.Encode()
	req.Header.Add("ContentType", JSON_HEADER)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return contacts, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return contacts, err
	}

	fmt.Println(string(b))
	return contacts, err
}
