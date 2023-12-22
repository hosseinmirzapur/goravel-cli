package utils

import (
	"log"
	"os"
	"strings"
)

func Error(cmdSection, desc string, err error) {
	if err != nil {
		log.Printf(
			"%s%s:%s/%s\n",
			Red,
			strings.ToUpper(cmdSection),
			desc,
			Reset,
		)
		log.Println("err: ", err)
		os.Exit(1)
	}
}

func Info(msg string) {
	log.Println(
		Blue,
		strings.Repeat("#", 10),
		msg,
		strings.Repeat("#", 10),
		Reset,
	)
}

func Success(msg string) {
	log.Println(
		Green,
		strings.Repeat("#", 10),
		msg,
		strings.Repeat("#", 10),
		Reset,
	)
}

func Alert(msg string) {
	log.Println(
		Yellow,
		strings.Repeat("#", 10),
		msg,
		strings.Repeat("#", 10),
		Reset,
	)
}
