package greetings

import (
	"errors"
	"fmt"
)



func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = msg
	}
	return messages, nil
}

func SayHi(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
