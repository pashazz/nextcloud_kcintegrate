package utils

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
)

func CheckUrl (s string, name string) error {
	_, err := url.Parse(s)
	if err != nil {
		return errors.New(fmt.Sprintf("%v: expected: valid URL; found: %v, error: %v", name, s, err))
	}
	return nil

}

func GetenvNonEmpty(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Fatalf("%v: not set", name)
	}
	return env
}
