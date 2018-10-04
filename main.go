package main

import (
	"fmt"
	"itchat4go/components"
)

func main() {
	uuid, err := components.GetQRuuid()
	components.CheckErr(err)
	fmt.Println(uuid)
}
