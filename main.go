package main

import (
	"fmt"
	c "itchat4go/components"
)

func main() {
	c.Login()
	fmt.Printf("Login succeed as %s", c.GetLoginUserName())
}
