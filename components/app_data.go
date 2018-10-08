package components

import (
	"net/http"
	"net/http/cookiejar"
)

type LoginData struct {
	baseReq  BaseRequest
	initInfo InitInfo
	info     map[string]string
}

var client http.Client
var loginData = LoginData{}

func init() {
	loginData.info = make(map[string]string)

	jar, _ := cookiejar.New(nil)
	client = http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Do not allow redirect
		},
	}
}
