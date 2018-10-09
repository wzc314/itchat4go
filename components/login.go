package components

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Login() {
	var uuid string
	for {
		fmt.Println("Requesting uuid of QR code.")
		if id, err := getQRuuid(); err != nil {
			fmt.Println("Requesting uuid failed. Try again.")
			PrintErr(err)
		} else {
			uuid = id
			break
		}
		SleepSec(1)
	}
	fmt.Println("Got the uuid of QR code: " + uuid)

	for {
		fmt.Println("Downloading QR code.")
		if err := getQR(uuid); err != nil {
			fmt.Println("Downloading QR code failed. Try again.")
			PrintErr(err)
		} else {
			break
		}
		SleepSec(1)
	}

	fmt.Println("Please scan the QR code to login.")

	for {
		fmt.Println("Checking the status of login.")
		if status, loginContent := checkLogin(uuid); status == 200 {
			err := processLoginInfo(loginContent)
			CheckErr(err)
			break
		}
		SleepSec(1)
	}

	webInit()
	showMobileLogin()
	GetContact()
	exec.Command("clear")
}

func getQRuuid() (string, error) {
	req, _ := http.NewRequest("GET", UUID_URL, nil)
	req.URL.RawQuery = url.Values{
		"appid": {"wx782c26e4c19acffb"},
		"fun":   {"new"},
		"lang":  {"zh_CN"},
		"_":     {GetTimestamp()}, // _: timestamp
	}.Encode()
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`^window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+)";$`)
	matches := re.FindStringSubmatch(string(b))
	status, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", err
	}

	if status != 200 {
		return "", errors.New(fmt.Sprintf("QR login code is %d", status))
	}

	return matches[2], nil
}

func getQR(uuid string) error {
	resp, err := http.Get(QRCODE_URL + uuid)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dst, err := os.Create(DEFAULT_QR)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, resp.Body)
	return err
}

func checkLogin(uuid string) (int, string) {
	var status, loginContent = 0, ""

	req, err := http.NewRequest("GET", CHECKLOGIN_URL, nil)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	req.URL.RawQuery = url.Values{
		"loginicon": {"true"},
		"uuid":      {uuid},
		"tip":       {"1"},
		"r":         {GetR()},
		"_":         {GetTimestamp()},
	}.Encode()
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	loginContent = string(b)

	re := regexp.MustCompile(`^window.code=(\d+);`)
	matches := re.FindStringSubmatch(loginContent)
	status, err = strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}

	switch status {
	case 200:
		fmt.Println("Login succeed initially.")
	case 201:
		fmt.Println("Please press confirm on your phone.")
	case 408:
		fmt.Println("Please scan the QR code.")
	default:
		fmt.Printf("Unknown return status code: %d\n", status)
	}

	return status, loginContent
}

func processLoginInfo(loginContent string) error {
	re := regexp.MustCompile(`window.redirect_uri="(\S+)";`)
	matches := re.FindStringSubmatch(loginContent)
	loginData.info["url"] = matches[1]

	req, _ := http.NewRequest("GET", loginData.info["url"], nil)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	result := LoginCallbackXMLResult{}
	err = xml.Unmarshal(b, &result)
	if err != nil {
		return err
	}

	loginData.baseReq.DeviceID = GetRandomID(15)
	loginData.baseReq.Skey = result.Skey
	loginData.baseReq.Sid = result.WXSid
	loginData.baseReq.Uin = result.WXUin
	loginData.info["pass_ticket"] = result.PassTicket
	loginData.info["logintime"] = GetTimestamp()
	loginData.info["url"] = loginData.info["url"][:strings.LastIndex(loginData.info["url"], "/")]

	indexUrl := req.URL.Hostname()
	if detailedUrl, ok := DetailedUrls[indexUrl]; ok {
		loginData.info["fileUrl"] = fmt.Sprintf("https://%s/cgi-bin/mmwebwx-bin", detailedUrl[0])
		loginData.info["syncUrl"] = fmt.Sprintf("https://%s/cgi-bin/mmwebwx-bin", detailedUrl[1])
	} else {
		loginData.info["fileUrl"] = loginData.info["url"]
		loginData.info["syncUrl"] = loginData.info["url"]
	}

	return nil
}

func webInit() {
	data := map[string]interface{}{}
	data["BaseRequest"] = loginData.baseReq

	b, err := json.Marshal(data)
	CheckErr(err)

	req, _ := http.NewRequest("POST", loginData.info["url"]+"/webwxinit", strings.NewReader(string(b)))
	req.URL.RawQuery = url.Values{
		"r":           {GetR()},
		"pass_ticket": {loginData.info["pass_ticket"]},
		"lang":        {"zh_CN"},
	}.Encode()
	req.Header.Add("ContentType", JSON_HEADER)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	CheckErr(err)

	initInfo := InitInfo{}
	err = json.Unmarshal(b, &initInfo)
	CheckErr(err)
	loginData.user = initInfo.User
	loginData.info["synckey"] = initInfo.SyncKey.ToString()
}

func showMobileLogin() {
	data := map[string]interface{}{}
	data["BaseRequest"] = loginData.baseReq
	data["Code"] = 3
	data["FromUserName"] = loginData.user.UserName
	data["ToUserName"] = loginData.user.UserName
	data["ClientMsgId"] = GetTimestamp()

	b, err := json.Marshal(data)
	CheckErr(err)

	req, _ := http.NewRequest("POST", loginData.info["url"]+"/webwxstatusnotify", strings.NewReader(string(b)))
	req.Header.Add("ContentType", JSON_HEADER)
	req.Header.Add("User-Agent", USER_AGENT)

	_, err = client.Do(req)
	PrintErr(err)
}
