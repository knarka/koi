package main

import (
	"github.com/nsf/termbox-go"

	"github.com/knarka/koi/koi"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	koi := koi.NewKoi()

	koi.Loop()
}
