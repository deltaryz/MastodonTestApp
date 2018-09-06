package main

import "fmt"

// DebugPrint will output caveman debug messages explaining exactly every step being taken in detail.
// It is recommended to call this function repeatedly for each line of text (so the "DEBUG:" text shows up properly.)
func DebugPrint(msg string) {
	// TODO: impl more graceful way to differentiate debug messages without needing to call the func repeatedly
	if config.DebugEnabled {
		fmt.Println("DEBUG: " + msg)
	}
}
