package toofast

import (
"fmt"
"net")

func main() {
	for i := 1; i < 1024; i++ {
	go func(j int) {
		address := fmt.Sprintf("scanme.nmap.org:%d", j)
	}
 }
}
