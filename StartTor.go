package main

import (
	"os/exec"
	"strconv"
	"log"
)

func main() {
	i := 10
	for i > 0{
		i--
		err := exec.Command("tor", "-socksport", strconv.Itoa(9000 + i)).Start()
		if err !=nil {
			log.Println(err)
			returnS
		}
	}
}
