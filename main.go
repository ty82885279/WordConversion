package main

import (
	"WordConversion/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Printf("cmd.Execute err:%v", err)
	}
}
