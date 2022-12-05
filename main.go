package main

import (
	"context"
	_ "context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	_ "log"
	"os"
	"strconv"
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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2327668012375-4462902086438-TEUjbSIliUjZvuZfX1VowKFJ")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04DQFKVA13-4493240691024-1597ceac4eeaadac62535f4993d7a9ced35c7be0e0380780ceeceb08c8662381")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My yop is <year>", &slacker.CommandDefinition{
		Description: "yop calculator",
		Example:     "My yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2021 - yob
			r := fmt.Sprint("age is %d", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
