package main

import (
	"fmt"
)

func SayHello(arg Args) {
	if len(arg.Strings) != 1 {
		fmt.Printf("Hello, !\n")
	} else {
		fmt.Printf("Hello, %s!\n", arg.Strings[0])
	}
}

func Quit(arg Args) {
	XGBConn.Quit = true
}

//func ToggleBar(arg Args) {
//	if BarVisible {
//		xproto.UnmapWindow(XGBConn, BarWindow)
//		BarVisible = false
//	} else {
//		xproto.MapWindow(XGBConn, BarWindow)
//		BarVisible = true
//	}
//}
