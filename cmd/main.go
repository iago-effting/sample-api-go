package main

import (
	"fmt"
	"iago-effting/go-template/pkg/http"
	"iago-effting/go-template/pkg/version"
	"log"
	"os"
)

func main() {
	fmt.Println("Version: ", version.Version)
	fmt.Println("Time Release:", version.Time)

	err := http.Run()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
