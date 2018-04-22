package configmodify

import (
	"log"
)

func FatalErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, ":", err)
	}
}

func PrintErr(msg string, err error) {
	if err != nil {
		log.Println(msg, ":", err)
	}
}
