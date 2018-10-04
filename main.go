package main

import (
	"fmt"
	c "itchat4go/components"
)

func main() {
	uuid, err := c.GetQRuuid()
	c.CheckErr(err)
	fmt.Println("Got the uuid of QR code: " + uuid)

	fmt.Println("Downloading QR code.")
	err = c.GetQR(uuid)
	c.CheckErr(err)

	fmt.Println("Please scan the QR code to log in.")
	fmt.Printf("%d", c.CheckLogin(uuid))
}
