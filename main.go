package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
    input := os.Args[1]
    fmt.Println(input)

    switch input {
    }

}

func createFile() {
    file, err := os.Create("tasks.json")
    if err != nil {
        log.Fatal("Can't create the file")
    }

    defer file.Close()
}
