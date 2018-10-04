package components

const (
	BASE_URL   = "https://login.weixin.qq.com"
	UUID_URL   = BASE_URL + "/jslogin"
	DEFAULT_QR = "QR.png"
	USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
)

var (
	UuidParams = map[string]string{
		"appid": "wx782c26e4c19acffb",
		"fun":   "new",
		"lang":  "zh_CN",
		"_":     "", // timestamp
	}
)
