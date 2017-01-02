package main

import (
	"fmt"
	"github.com/OSPG/inproxy"
)

func main() {
	proxy := inproxy.NewProxy(":8080", inproxy.DebugLevel)
	err := proxy.Serve()
	if err != nil {
		fmt.Println(err)
	}
	for {
	}
}
