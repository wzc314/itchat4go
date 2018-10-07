package components

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
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

func PrintErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func getTimestamp() int {
	return int(time.Now().UnixNano() / 1000000)
}

func GetTimestamp() string {
	return strconv.Itoa(getTimestamp())
}

func GetR() string {
	return strconv.Itoa(-getTimestamp() / 1579)
}

func SleepSec(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func GetRandomID(n int) string {
	rand.Seed(time.Now().Unix())
	return "e" + strconv.FormatFloat(rand.Float64(), 'f', n, 64)[2:]
}
