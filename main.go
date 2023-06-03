package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

func main() {
	log.Println("Opening rpio")
	err := rpio.Open() // Open up our access to the GPIO pins (requires sudo)
	if err != nil {
		log.Panicf("could not open rpio: %s", err)
	}
	defer rpio.Close()

	relay1Pin := rpio.Pin(18)
	relay1Pin.Output()
	relay1Pin.High()

}
