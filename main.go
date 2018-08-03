package main

import (
	"fmt"

	"github.com/goware/saml"
	"github.com/mfridman/srfax"
)

var securityOpts = &saml.SecurityOpts{}

func main() {
	// contrived example
	c := getClient()
	ok, _ := c.CheckAuth()
	fmt.Println(ok)
	fmt.Println(securityOpts)
}

func getClient() *srfax.Client {
	c, _ := srfax.NewClient(srfax.ClientCfg{ID: 00001, Pwd: "password"})
	return c
}
