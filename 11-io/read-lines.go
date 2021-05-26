package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("sample.txt")
	if fileError != nil {
		log.Fatalln(fileError)
	}
	//defer fileHandle.Close()
	inputReader := bufio.NewReader(fileHandle)
	for {
		line, err := inputReader.ReadString('.')
		fmt.Println(line)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
