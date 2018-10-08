package components

import (
	"encoding/xml"
	"fmt"
)

type LoginCallbackXMLResult struct {
	XMLName     xml.Name `xml:"error"`
	Ret         string   `xml:"ret"`
	Message     string   `xml:"message"`
	Skey        string   `xml:"Skey"`
	WXSid       string   `xml:"wxsid"`
	WXUin       string   `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	IsGrayscale string   `xml:"isgrayscale"`
}

type BaseRequest struct {
	Uin      string `json:"Uin"`
	Sid      string `json:"Sid"`
	Skey     string `json:"Skey"`
	DeviceID string `json:"DeviceID"`
}

type InitInfo struct {
	User    User    `json:"User"`
	SyncKey SyncKey `json:"SyncKey"`
}

type User struct {
	Uin        int64  `json:"Uin"`
	UserName   string `json:"UserName"`
	NickName   string `json:"NickName"`
	RemarkName string `json:"RemarkName"`
	Sex        int8   `json:"Sex"`
	Signature  string `json:"Signature"`
}

type SyncKey struct {
	Count int  `json:"Count"`
	List  []KV `json:"List"`
}

type KV struct {
	Key int64 `json:"Key"`
	Val int64 `json:"Val"`
}

func (sk SyncKey) ToString() string {
	var s string
	for i := 0; i < sk.Count; i++ {
		s = s + fmt.Sprintf("%d_%d|", sk.List[i].Key, sk.List[i].Val)
	}
	return s[:len(s)-1]
}
