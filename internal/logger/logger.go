package logger

import (
	"encoding/json"
	"log"
)

func Struct(o any) {
	bytes, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(bytes))
	}
}
