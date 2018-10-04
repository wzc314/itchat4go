package components

import (
	"net/url"
	"time"
)

func GetParams(data map[string]string) string {
	params := url.Values{}
	for k, v := range data {
		params.Add(k, v)
	}
	return params.Encode()
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
