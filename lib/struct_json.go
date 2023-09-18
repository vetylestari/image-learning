package lib

import (
	"encoding/json"
	"log"
)

func StructToJSON(structt any) string {
	json_byte, err := json.Marshal(structt)
	if err != nil {
		log.Fatal("Error StructToJSON")
	}
	return string(json_byte)
}
