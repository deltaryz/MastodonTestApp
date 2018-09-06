package main

import (
	"context"
	"fmt"
	"github.com/mattn/go-mastodon"
	"log"
)

// DebugPrint will output caveman debug messages explaining exactly every step being taken in detail.
// It is recommended to call this function repeatedly for each line of text (so the "DEBUG:" text shows up properly.)
func DebugPrint(message string) {
	// TODO: impl more graceful way to differentiate debug messages without needing to call the func repeatedly
	if config.DebugEnabled {
		fmt.Println("DEBUG: " + message)
	}
}

// RetrieveHomeTimeline will retrieve the toots in your Home timeline, returning []*Status.
func RetrieveHomeTimeline() []*mastodon.Status {
	// retrieve timeline
	DebugPrint("Now retrieving home timeline")
	timeline, err := c.GetTimelineHome(context.Background(), nil)
	if err != nil {
		fmt.Println("Something went wrong with retrieving the timeline:\n")
		log.Fatal(err)
	}
	return timeline
}

// PostToot will post a new status message and return the status object.
func PostToot(message string) (*mastodon.Status, error) {
	// post a test message
	DebugPrint("Now posting a message")
	DebugPrint("Message text: " + message)
	status, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status: "This is a test message, don't mind me.",
	})

	return status, err
}
