package main

import (
	//"encoding/json"
	//"fmt"
	"log"
	"os"
)

func main() {
    createFile()


}

func createFile() {
    file, err := os.Create("tasks.json")
    if err != nil {
        log.Fatal("Can't create the file")
    }

    defer file.Close()
}
