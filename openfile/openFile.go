package openfile

import (
	"log"
	"os"
)

func Fileopen(filename string, txtfile chan <- string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("'test1.txt' faylni o'qishda xatolik", err)
	}
	txtfile <- string(data)
}