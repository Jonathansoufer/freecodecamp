package main

import (
	"fmt"
	"log"
	"os"
	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	channelArr := []string{os.Getenv("SLACK_CHANNEL")}

	fileArr := []string{os.Getenv("FILE_PATH")}

	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			File: fileArr[i],
			Channels: channelArr,
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
	
}