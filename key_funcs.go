package main

import (
	"fmt"

	"github.com/jezek/xgb/xproto"
)

func SayHello(arg Args) {
	fmt.Printf("Hello, %s", arg.Strings[0])
}

func Quit(arg Args) {
	RuntimeFlag = false
}

func ToggleBar(arg Args) {
	if BarVisible {
		xproto.UnmapWindow(XGBConn, BarWindow)
		BarVisible = false
	} else {
		xproto.MapWindow(XGBConn, BarWindow)
		BarVisible = true
	}
}
