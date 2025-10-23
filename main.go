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
var CurrentKeyMask BitFlags

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

	for i := 0; i < len(Hotkeys); i++ {
		Hotkeys[i].InitMask()
	}
	xproto.GrabKeyboard(XGBConn, true, root,
		xproto.TimeCurrentTime,
		xproto.GrabModeAsync, xproto.GrabModeAsync)

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
			// 00000001|10000000

			CurrentKeyMask.Set(e.Detail)

			for _, hk := range Hotkeys {
				if CurrentKeyMask.And(hk.KeyMask) == hk.KeyMask {
					hk.Action(hk.Arguments)
				}
			}

			//for _, hk := range Hotkeys {
			//	if hk.Key == 0 && e.State&hk.Modifiers == hk.Modifiers {
			//		hk.Action(hk.Arguments)
			//	} else if e.Detail == xproto.Keycode(hk.Key) && e.State&hk.Modifiers == hk.Modifiers {
			//		hk.Action(hk.Arguments)
			//	}
			//}
		case xproto.KeyReleaseEvent:
			//fmt.Printf("Key released: keycode=%d\n", e.Detail)
			CurrentKeyMask.Clear(e.Detail)
		}
	}
}
