package main

import (
	"fmt"
	"strings"
)

func main() {
	message := " 05:38:52 up  4:29,  3 users,  load average: 1.52, 1.77, 1.81"

	data := strings.Split(message, " ")

	for key, values := range data {
		if values != "" {
			fmt.Println(key, values)
		}
	}
}
