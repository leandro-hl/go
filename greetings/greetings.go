package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	mge := fmt.Sprintf(randomFormat(), name)

	return mge, nil
}

func Hellos(names []string) (map[string]string, []error) {
	mges := make(map[string]string)
	errors := []error{}

	for _, name := range names {
		mge, err := Hello(name)

		if err != nil {
			errors = append(errors, err)
		} else {
			mges[name] = mge
		}
	}

	return mges, errors
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v well met!",
	}

	return formats[rand.Intn(len(formats))]
}
