package components

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func GetQRuuid() (string, error) {
	client := http.Client{}

	UuidParams["_"] = strconv.Itoa(int(time.Now().Unix()))
	req, err := http.NewRequest("GET", UUID_URL, nil)
	CheckErr(err)
	req.URL.RawQuery = GetParams(UuidParams)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	reg := regexp.MustCompile(`^window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+)";$`)
	matches := reg.FindStringSubmatch(string(bodyBytes))
	status, err := strconv.Atoi(matches[1])
	CheckErr(err)

	if status != 200 {
		return "", errors.New(fmt.Sprintf("QR login code is %d", status))
	}

	return matches[2], nil
}
