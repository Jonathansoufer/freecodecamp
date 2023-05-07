package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	slack_bot_token := os.Getenv("SLACK_BOT_TOKEN")
	slack_app_token := os.Getenv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(slack_bot_token, slack_app_token)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("hello", &slacker.CommandDefinition{
		Description: "Say Hello",
		Example:     "hello",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}