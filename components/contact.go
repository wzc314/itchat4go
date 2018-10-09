package components

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetContact() error {
	req, _ := http.NewRequest("GET", loginData.info["url"]+"/webwxgetcontact", nil)
	req.URL.RawQuery = url.Values{
		"lang":        {"zh_CN"},
		"pass_ticket": {loginData.info["pass_ticket"]},
		"r":           {GetTimestamp()},
		"seq":         {"0"},
		"skey":        {loginData.baseReq.Skey},
	}.Encode()
	req.Header.Add("ContentType", JSON_HEADER)
	req.Header.Add("User-Agent", USER_AGENT)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	memberList := MemberList{}
	err = json.Unmarshal(b, &memberList)
	CheckErr(err)

	for i := 0; i < memberList.MemberCount; i++ {
		contacts[memberList.Users[i].UserName] = memberList.Users[i]
		if strings.HasPrefix(memberList.Users[i].UserName, "@@") {
			chatroomList[memberList.Users[i].UserName] = memberList.Users[i]
		}
	}

	for _, member := range contacts {
		fmt.Println(member.NickName)
	}
	return nil
}
