package main

import "github.com/schweigert/mmorpg-communication-analysis/infrastructure/web"

func main() {
	web.NewWServer().Start()
}
