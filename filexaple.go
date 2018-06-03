package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tmpfile, err := os.Open("main.go")
	if err != nil {
		log.Fatal("%v", err)

	}
	defer tmpfile.Close()
	fileinfo, _ := tmpfile.Stat()
	fmt.Print(fileinfo.Size())
	buff := make([]byte, 100)
	size, err := tmpfile.ReadAt(buff, 2197)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x", buff[:size])
	fmt.Print("content is ", string(buff[:size]))

}
