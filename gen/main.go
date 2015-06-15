package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "message_lookup":
		genMessageLookup()
	case "game_events":
		genGameEventLookup()
	default:
		fmt.Println("unknown mode", os.Args[1])
	}
}
