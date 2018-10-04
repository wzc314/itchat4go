package components

const (
	BASE_URL       = "https://login.weixin.qq.com"
	UUID_URL       = BASE_URL + "/jslogin"
	QRCODE_URL     = BASE_URL + "/qrcode/"
	CHECKLOGIN_URL = BASE_URL + "/cgi-bin/mmwebwx-bin/login"
	DEFAULT_QR     = "QR.jpg"
	USER_AGENT     = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
)

var (
	UuidParams = map[string]string{
		"appid": "wx782c26e4c19acffb",
		"fun":   "new",
		"lang":  "zh_CN",
		"_":     "", // timestamp
	}

	CheckLoginParams = map[string]string{
		"loginicon": "true",
		"uuid":      "",
		"tip":       "1",
		"r":         "",
		"_":         "",
	}
)
