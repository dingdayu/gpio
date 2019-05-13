package main

import (
	"fmt"
	"gpio"
	"log"
	"time"
)

func main() {
	_ = gpio.New()

	p := gpio.NewPin(18, gpio.OUT)
	err := p.Export()
	log.Println(err)
	for {
		err = p.Write(1)
		time.Sleep(5 * time.Second)
		err = p.Write(0)
		time.Sleep(5 * time.Second)
		fmt.Println("……")
	}
}
