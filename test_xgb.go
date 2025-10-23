package main

import (
	"fmt"
	"log"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func test_xgb() {
	// Подключаемся к X серверу
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Не удалось подключиться к X серверу: %v", err)
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)

	fmt.Printf("Подключено к X серверу!\n")
	fmt.Printf("Root window: %d\n", screen.Root)
	fmt.Printf("Размер экрана: %dx%d\n", screen.WidthInPixels, screen.HeightInPixels)
}
