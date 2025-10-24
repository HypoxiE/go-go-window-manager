package main

import (
	"fmt"
	"log"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var BarWindow xproto.Window
var XGBConn *xgb.Conn
var RuntimeFlag = true

var masks = []uint16{
	0,
	xproto.ModMaskLock, // CapsLock
	xproto.ModMask2,    // NumLock
	xproto.ModMaskLock | xproto.ModMask2,
}

func main() {
	var err error

	XGBConn, err = xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer XGBConn.Close()

	setup := xproto.Setup(XGBConn)
	screen := setup.DefaultScreen(XGBConn)
	root := screen.Root

	BarWindow, _ = xproto.NewWindowId(XGBConn)
	width := uint16(screen.WidthInPixels)
	height := uint16(BarHeight)

	xproto.CreateWindow(XGBConn, screen.RootDepth, BarWindow, root,
		0, 0, width, height, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{BarBackground, xproto.EventMaskExposure | xproto.EventMaskButtonPress})

	xproto.ChangeWindowAttributes(XGBConn, BarWindow, xproto.CwEventMask,
		[]uint32{xproto.EventMaskExposure | xproto.EventMaskButtonPress |
			xproto.EventMaskKeyPress | xproto.EventMaskKeyRelease})

	for _, hk := range Hotkeys {
		for _, m := range masks {
			xproto.GrabKey(XGBConn, false, root, hk.Modifier|m, hk.Key, xproto.GrabModeAsync, xproto.GrabModeAsync)
		}
	}

	BarVisible = !BarVisible
	ToggleBar(Args{})

	for RuntimeFlag {
		ev, err := XGBConn.WaitForEvent()
		if err != nil {
			log.Fatal(err)
		}

		switch e := ev.(type) {
		//case xproto.ExposeEvent:
		//	fmt.Printf("Окно отрисовано")

		case xproto.KeyPressEvent:
			fmt.Printf("Key pressed: keycode=%d\n", e.Detail)

			for _, hk := range Hotkeys {
				if hk.Key == 0 && e.State&hk.Modifier == hk.Modifier {
					hk.Action(hk.Arguments)
				} else if e.Detail == xproto.Keycode(hk.Key) && e.State&hk.Modifier == hk.Modifier {
					hk.Action(hk.Arguments)
				}
			}
		}
	}
}
