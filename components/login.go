package components

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
	"strconv"
)

var client http.Client

func init() {
	client = http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	} // Do not allow redirect
	jar, err := cookiejar.New(nil)
	CheckErr(err)
	client.Jar = jar
}

func Login() {
	uuid, err := getQRuuid()
	CheckErr(err)
	fmt.Println("Got the uuid of QR code: " + uuid)

	fmt.Println("Downloading QR code.")
	err = getQR(uuid)
	CheckErr(err)
	fmt.Println("Please scan the QR code to login.")

	for {
		fmt.Println("Checking the status of login.")

		if status, loginContent := checkLogin(uuid); status == 200 {
			processLoginInfo(loginContent)
			break
		}
		SleepSec(1)
	}
}

func getQRuuid() (string, error) {
	UuidParams["_"] = strconv.Itoa(int(GetTimestamp()))
	req, err := http.NewRequest("GET", UUID_URL, nil)
	CheckErr(err)
	req.URL.RawQuery = GetParams(UuidParams)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	re := regexp.MustCompile(`^window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+)";$`)
	matches := re.FindStringSubmatch(string(bodyBytes))
	status, err := strconv.Atoi(matches[1])
	CheckErr(err)

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

	CheckLoginParams["uuid"] = uuid
	CheckLoginParams["_"] = strconv.Itoa(int(GetTimestamp()))
	CheckLoginParams["r"] = strconv.Itoa(int(-GetTimestamp() / 1579))

	req, err := http.NewRequest("GET", CHECKLOGIN_URL, nil)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	req.URL.RawQuery = GetParams(CheckLoginParams)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}
	loginContent = string(bodyBytes)

	re := regexp.MustCompile(`^window.code=(\d+);`)
	matches := re.FindStringSubmatch(loginContent)
	status, err = strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(err)
		return status, loginContent
	}

	switch status {
	case 200:
		fmt.Println("Log in succeed.")
	case 201:
		fmt.Println("Please press confirm on your phone.")
	case 408:
		fmt.Println("Please scan the QR code.")
	default:
		fmt.Printf("Unknown return status code: %d\n", status)
	}

	return status, loginContent
}

func processLoginInfo(loginContent string) {
	re := regexp.MustCompile(`window.redirect_uri="(\S+)";`)
	matches := re.FindStringSubmatch(loginContent)
	LoginInfo["url"] = matches[1]

	req, err := http.NewRequest("GET", LoginInfo["url"], nil)
	CheckErr(err)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}
