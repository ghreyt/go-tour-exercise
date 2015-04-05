package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
