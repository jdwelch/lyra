package mock

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"

	"github.com/mgutz/ansi"
)

// MessageType is a Type for the 'enum' below
type MessageType int

// Kinds of messages
const (
	Question = iota
	Process
	Banner
	Plain
)

// Exists checks to see if a file or dir exists
func Exists(filePath string) (exists bool) {
	exists = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exists = false
	}
	return
}

// SetDirty writes value to Viper
func SetDirty() {
	viper.Set("apply-dirty", true)
	viper.WriteConfig()
}

// UnsetDirty removes ^
func UnsetDirty() {
	viper.Set("apply-dirty", false)
	viper.WriteConfig()
}

// Delay artifically slows down execution
func Delay(durationMs int) {
	time.Sleep(time.Duration(durationMs) * time.Millisecond)
}

// Outline is a shorthand for sketching out the flow of a command
func Outline(kind MessageType, message string) {
	switch kind {
	case Process:
		fmt.Println(ansi.Blue + message + ansi.Reset)
	case Question:
		fmt.Println(ansi.Green + "? " + ansi.Reset + message + " [y/n]")
	case Banner:
		fmt.Println(ansi.Cyan + "==== " + message + " ====" + ansi.Reset)
	case Plain:
		fmt.Println(message)
	default:
		fmt.Println(message)
	}
}

// OutlineSurvey fakes the multiple-choice part of AlecAivazis/survey in a super simple and stupid way, hooray!
func OutlineSurvey(question string, options []string) {
	fmt.Println(ansi.Green + "? " + ansi.Reset + question)
	for _, v := range options {
		fmt.Println(ansi.Cyan + "  ❯ " + ansi.Reset + v)
	}
}
