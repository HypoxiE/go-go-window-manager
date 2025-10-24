package main

import (
	"log"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

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

	DrawBar()
	DrawCursor()

	for _, hk := range Hotkeys {
		for _, m := range masks {
			xproto.GrabKey(XGBConn, false, root, hk.Modifier|m, hk.Key, xproto.GrabModeAsync, xproto.GrabModeAsync)
		}
	}

	for RuntimeFlag {
		ev, err := XGBConn.WaitForEvent()
		if err != nil {
			log.Fatal(err)
		}

		switch e := ev.(type) {
		//case xproto.ExposeEvent:
		//	fmt.Printf("Окно отрисовано")

		case xproto.KeyPressEvent:
			//fmt.Printf("Key pressed: keycode=%d\n", e.Detail)

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
