package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting.
// In Go, a function whose name starts with a capital letter
// can be called by a function not in the same package.
// This is known in Go as an exported name.
func Hello(name string) (string, error) {
	// Handle no name param passed
	if name == "" {
		return "", errors.New("'name' can't be empty")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())	// break the function to confirm tests
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map top associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function for each of them.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Grate to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected msg by speccing a random index val
	return formats[rand.Intn(len(formats))]
}
