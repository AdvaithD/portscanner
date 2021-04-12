package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(address)
		conn, err := net.Dial("tcp", "scanme.nmap.org:80")
		if err != nil {
			fmt.Println("Error connecting")
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
