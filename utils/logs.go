package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Error(cmdSection, desc string, err error) {
	if err != nil {
		fmt.Printf(
			"%s%s: %s%s\n",
			Red,
			strings.ToUpper(cmdSection),
			desc,
			Reset,
		)
		log.Println("err: ", err)
		os.Exit(1)
	}
}

func Info(msg string, separated bool) {
	if separated {
		log.Println(separate(msg, Blue))
	} else {
		log.Println(Blue, msg, Reset)
	}

}

func Success(msg string, separated bool) {
	if separated {
		log.Println(separate(msg, Green))
	} else {
		log.Println(Green, msg, Reset)
	}
}

func Alert(msg string, separated bool) {
	if separated {
		log.Println(separate(msg, Yellow))
	} else {
		log.Println(Yellow, msg, Reset)
	}
}

func separate(msg, color string) string {
	return fmt.Sprintf(
		"%s%s%s%s%s",
		color,
		strings.Repeat("#", 10),
		msg,
		strings.Repeat("#", 10),
		Reset,
	)

}
