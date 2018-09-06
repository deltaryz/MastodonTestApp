package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/mattn/go-mastodon"
	"log"
	"strconv"
	"time"
)

// Environment variables
type environmentVariables struct {
	ClientID         string `env:"CLIENT_ID" envDefault:""`          // environment variable CLIENT_ID
	ClientSecret     string `env:"CLIENT_SECRET" envDefault:""`      // environment variable CLIENT_SECRET
	MastodonServer   string `env:"MASTODON_SERVER" envDefault:""`    // environment variable MASTODON_SERVER
	MastodonEmail    string `env:"MASTODON_EMAIL" envDefault:""`     // environment variable MASTODON_EMAIL
	MastodonPassword string `env:"MASTODON_PASSWORD" envDefault:""`  // environment variable MASTODON_PASSWORD
	DebugEnabled     bool   `env:"DEBUG_ENABLED" envDefault:"false"` // environment variable DEBUG_ENABLED
}

// Global variables
var (
	config environmentVariables
)

func main() {
	// init environment variables
	DebugPrint("Now processing environment variables")
	config = environmentVariables{}
	configErr := env.Parse(&config)
	if configErr != nil {
		fmt.Println("Error processing environment variables.\nPlease check https://github.com/techniponi/MastodonTestApp for details.\n\n" + configErr.Error())
		return
	}
	// initiate client
	DebugPrint("Now initiating client with the following server, ID and secret")
	DebugPrint(config.MastodonServer)
	DebugPrint(config.ClientID)
	DebugPrint(config.ClientSecret)
	c := mastodon.NewClient(&mastodon.Config{
		Server:       config.MastodonServer,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	})

	// authenticate
	DebugPrint("Using login:")
	DebugPrint(config.MastodonEmail)
	DebugPrint(config.MastodonPassword)
	DebugPrint("Now attempting to authenticate")
	err := c.Authenticate(context.Background(), config.MastodonEmail, config.MastodonPassword)
	if err != nil {
		fmt.Println("Something went wrong with authentication:\n")
		log.Fatal(err)
	}

	/*
		// retrieve timeline
		DebugPrint("Now retrieving timeline")
		timeline, err := c.GetTimelineHome(context.Background(), nil)
		if err != nil {
			fmt.Println("Something went wrong with retrieving the timeline:\n")
			log.Fatal(err)
		}

		// display timeline
		DebugPrint("Now displaying timeline")
		for i := len(timeline) - 1; i >= 0; i-- {
			fmt.Println(timeline[i])
		}
	*/

	// post a test message
	DebugPrint("Now posting test message:")
	status, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status: "This is a test message, don't mind me.",
	})

	fmt.Println(status)

	testMessageID := status.URL[len(status.URL)-18:]
	fmt.Println("This toot's ID is: " + testMessageID)

	DebugPrint("Now waiting 5 seconds before deleting...")
	time.Sleep(5 * time.Second)

	DebugPrint("Attempting to delete toot")
	testMessageIDAsInt, err := strconv.ParseInt(testMessageID, 10, 64)

	c.DeleteStatus(context.Background(), testMessageIDAsInt)
}
