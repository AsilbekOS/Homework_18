package main

import (
	"fmt"
	"homework_18/faktarial"
	"homework_18/openfile"
	"homework_18/openjsonfile"
)

func main() {
	txtfilename := "test1.txt"
	jsonfilename := "test2.json"

	factorial := make(chan int)
	txtfile := make(chan string)
	jsonfile := make(chan openjsonfile.User)

	go func() {
		factorial <- faktarial.Fact(<-factorial)
	}()
	var n int
	fmt.Println("Qaysi sonni faktorialini hisoblaymiz:")
	fmt.Print("Son:")
	fmt.Scan(&n)
	factorial <- n
	result := <-factorial
	fmt.Println(n, "- faktoriali =", result)

	go openfile.Fileopen(txtfilename, txtfile)
	data := <-txtfile
	fmt.Println("test1.txt filedagi ma'lumotlar:", data)

	go openjsonfile.Decoder(jsonfilename, jsonfile)
	fmt.Println("test2.json filedagi ma'lumotlar:", <-jsonfile)
	fmt.Println("test2.json filedagi ma'lumotlar:", <-jsonfile)
}
