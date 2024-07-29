package helper

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Configure(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		fmt.Printf("error at loading %s file", filename)
		return err
	}

	return nil
}
