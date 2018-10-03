package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"itchat4go/config"
	"itchat4go/utils"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	uuid, err := getQRuuid()
	utils.CheckErr(err)
	fmt.Println(uuid)
}

func getQRuuid() (string, error) {
	client := http.Client{}

	config.UuidParams["_"] = strconv.Itoa(int(time.Now().Unix()))
	req, err := http.NewRequest("GET", config.UUID_URL, nil)
	utils.CheckErr(err)
	req.URL.RawQuery = utils.GetParams(config.UuidParams)
	req.Header.Add("User-Agent", config.USER_AGENT)

	resp, err := client.Do(req)
	utils.CheckErr(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr(err)

	reg := regexp.MustCompile(`^window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+)";$`)
	matches := reg.FindStringSubmatch(string(bodyBytes))
	status, err := strconv.ParseInt(matches[1], 10, 64)
	utils.CheckErr(err)

	if status != 200 {
		return "", errors.New(fmt.Sprintf("QR login code is ", status))
	}

	return matches[2], nil
}
