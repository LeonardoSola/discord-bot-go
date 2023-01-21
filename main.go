package main

import (
	"fmt"
	"os"
	"os/signal"
	"discord-bot-go/config"
	"discord-bot-go/handlers"
	"syscall"
)

func main() {
	err := config.EnvSetup()
	if err != nil {
		fmt.Println("Environment: ❌")
		panic(err)
	}
	fmt.Println("Environment: ✅")

	err = config.DiscordSetup()
	if err != nil {
		fmt.Println("Discord: ❌")
		panic(err)
	}
	fmt.Println("Discord: ✅")

	handlers.SetupHandlers()

	err = config.DiscordOpen()
	if err != nil {
		fmt.Println("Session: ❌")
		panic(err)
	}
	fmt.Println("Session: ✅")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	config.DiscordClose()

}
