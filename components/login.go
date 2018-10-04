package components

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func GetQRuuid() (string, error) {
	client := http.Client{}

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

func GetQR(uuid string) error {
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

func CheckLogin(uuid string) int {
	client := http.Client{}

	CheckLoginParams["uuid"] = uuid
	CheckLoginParams["_"] = strconv.Itoa(int(GetTimestamp()))
	CheckLoginParams["r"] = strconv.Itoa(int(-GetTimestamp() / 1579))

	req, err := http.NewRequest("GET", CHECKLOGIN_URL, nil)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	req.URL.RawQuery = GetParams(CheckLoginParams)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	re := regexp.MustCompile(`^window.code=(\d+);`)
	matches := re.FindStringSubmatch(string(bodyBytes))
	status, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return status
}
