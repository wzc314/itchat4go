package components

import (
	"net/http"
	"net/http/cookiejar"
)

type LoginData struct {
	baseReq BaseRequest
	user    User
	info    map[string]string
}

var (
	client       http.Client
	loginData    = LoginData{}
	contacts     = make(map[string]User)
	chatroomList = make(map[string]User)
)

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

func GetLoginUserName() string {
	return loginData.user.NickName
}
