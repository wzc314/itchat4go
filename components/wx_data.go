package components

import "encoding/xml"

type LoginCallbackXMLResult struct {
	XMLName     xml.Name `xml:"error"`
	Ret         string   `xml:"ret"`
	Message     string   `xml:"message"`
	SKey        string   `xml:"skey"`
	WXSid       string   `xml:"wxsid"`
	WXUin       string   `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	IsGrayscale string   `xml:"isgrayscale"`
}

type BaseRequest struct {
	Uin      string `json:"Uin"`
	Sid      string `json:"Sid"`
	SKey     string `json:"Skey"`
	DeviceID string `json:"DeviceID"`
}
