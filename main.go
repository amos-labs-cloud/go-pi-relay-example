package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

func main() {
	log.Println("Opening rpio")
	err := rpio.Open() // Open access to the GPIO
	if err != nil {
		log.Panicf("could not open rpio: %s", err)
	}
	defer rpio.Close() // Close the connection once we are done with this function

	relay1Pin := 18 // This is the pin that we connected to our first relay
	log.Printf("connecting to pin %d\n", relay1Pin)
	relay1 := rpio.Pin(18) // instantiate the pin

	state := relay1.Read() // Read its current status
	if state == rpio.Low {
		log.Println("relay is currently 'off'")
	} else {
		log.Println("relay is currently 'on'")
	}

	relay1.Output() // Prepare the pin to be written to
	log.Println("toggling the relay")
	relay1.Toggle() // Toggle the relay, it switches from high to low and low to high
}
