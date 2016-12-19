package main

import (
//	"swap/protocol"
	"swap/modem"
	"log"
	"os"
)

func main() {
	modem := modem.SerialModem{Filename: os.Args[1]}

	log.Printf("Testing %s\n", modem.Filename)
	modem.Open()
	modem.Send("AT?\n")
	for done := false; !done; {
		line, err := modem.ReadLine()
		log.Printf("%s\n", line)
		if err != nil {
			log.Fatal(err)
		}
	}
}
