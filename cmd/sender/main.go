package main

import (
	"flag"
	"github.com/wmrodrigues/twitter-sender/internal/services/loader"
	"github.com/wmrodrigues/twitter-sender/internal/services/sender"
	"github.com/wmrodrigues/twitter-sender/internal/services/settings"
	"log"
)

func main() {
	log.Println("starting the twitter-sender...")

	log.Println("loading settings file")
	_settings := settings.LoadSettingsFile()

	var path string
	flag.StringVar(&path, "file", "", "specify the complete csv file path")
	flag.Parse()

	if path == "" {
		log.Fatal("file parameter is required, please call like this: ./mailer --file=some/existent/path/file.csv")
	}

	log.Printf("loading csv file from %s\n", path)
	recipients, err := loader.LoadFromCsvFile(path)

	if err != nil {
		log.Fatal(err)
	}

	_sender := sender.NewTwitter(_settings)
	_sender.SetRecipients(recipients)

	log.Println("now the fun part, sending twits")
	_sender.SendTweets()

	log.Println("sending process finished!")
}