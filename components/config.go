package components

const (
	BASE_URL       = "https://login.weixin.qq.com"
	UUID_URL       = BASE_URL + "/jslogin"
	QRCODE_URL     = BASE_URL + "/qrcode/"
	CHECKLOGIN_URL = BASE_URL + "/cgi-bin/mmwebwx-bin/login"
	DEFAULT_QR     = "QR.jpg"
	USER_AGENT     = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
	JSON_HEADER    = "application/json;charset=UTF-8"
)

var (
	DetailedUrls = map[string][2]string{
		"wx2.qq.com":      {"file.wx2.qq.com", "webpush.wx2.qq.com"},
		"wx8.qq.com":      {"file.wx8.qq.com", "webpush.wx8.qq.com"},
		"qq.com":          {"file.wx.qq.com", "webpush.wx.qq.com"},
		"web2.wechat.com": {"file.web2.wechat.com", "webpush.web2.wechat.com"},
		"wechat.com":      {"file.web.wechat.com", "webpush.web.wechat.com"},
	}
)
