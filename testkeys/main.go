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

	// set LED line off
	var ll [8]*gpiod.Line
	for k, l := range ledRow {
		ll[k], err = c.RequestLine(l, gpiod.AsOutput(0))
		if err != nil {
			panic(err)
		}
	}

	// set ROWs 0v (sinks current)
	var kr [3]*gpiod.Line
	for k, r := range rows {
		kr[k], err = c.RequestLine(r, gpiod.AsOutput(0))
		if err != nil {
			panic(err)
		}
		defer kr[k].Close()
	}

	// set columns as input
	period := 10 * time.Millisecond
	var kc [12]*gpiod.Line
	for k, col := range cols {
		kc[k], err = c.RequestLine(
			col,
			gpiod.WithPullUp,
			gpiod.WithBothEdges,
			gpiod.WithDebounce(period),
			gpiod.WithEventHandler(eventHandler))
		if err != nil {
			panic(err)
		}
		defer kc[k].Close()
	}

	<-time.After(40 * time.Second)
}

func eventHandler(evt gpiod.LineEvent) {
	t := time.Now()
	edge := "rising"
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
	}
	fmt.Printf("event:%3d %-7s %s (%s)\n",
		evt.Offset,
		edge,
		t.Format(time.RFC3339Nano),
		evt.Timestamp)
}
