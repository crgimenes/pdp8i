package main

import (
	"fmt"
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

var (
	cols = [12]int{
		rpi.GPIO14,
		rpi.GPIO15,
		rpi.GPIO4,
		rpi.GPIO5,
		rpi.GPIO6,
		rpi.GPIO7,
		rpi.GPIO8,
		rpi.GPIO9,
		rpi.GPIO10,
		rpi.GPIO11,
		rpi.GPIO12,
		rpi.GPIO13,
	}

	rows = [3]int{
		rpi.GPIO16,
		rpi.GPIO17,
		rpi.GPIO18,
	}

	ledRow = [8]int{
		rpi.GPIO20,
		rpi.GPIO21,
		rpi.GPIO22,
		rpi.GPIO23,
		rpi.GPIO24,
		rpi.GPIO25,
		rpi.GPIO26,
		rpi.GPIO27,
	}
)

func main() {
	fmt.Println("test")

	c, err := gpiod.NewChip("gpiochip0")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// set leader line on
	var ll [8]*gpiod.Line
	for k, l := range ledRow {
		ll[k], err = c.RequestLine(l, gpiod.AsOutput(1))
		if err != nil {
			panic(err)
		}
	}

	// set LEDs on
	var lc [12]*gpiod.Line
	for k, col := range cols {
		lc[k], err = c.RequestLine(col, gpiod.AsOutput(0))
		if err != nil {
			panic(err)
		}
	}

	// wait 3 seconds
	<-time.After(3 * time.Second)

	// set LEDs off
	for _, l := range lc {
		err = l.SetValue(1)
		if err != nil {
			panic(err)
		}
	}

}
