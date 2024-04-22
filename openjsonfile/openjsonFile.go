package openjsonfile

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func Decoder(filename string, jsonfile chan<- User) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var user []User
	if err := json.Unmarshal(file, &user); err != nil {
		log.Fatal("'test2.json' faylni o'qishda xatolik:", err)
	}

	for _, i := range user {
		jsonfile <- i
	}
}
