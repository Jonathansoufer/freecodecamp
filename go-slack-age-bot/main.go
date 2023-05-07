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

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Example:     "my yob is 1981",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.ReportError(err)
				return
			}
			age := time.Now().Year() - yob
			response.Reply(fmt.Sprintf("Your year of birth is %s", year))
			response.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}