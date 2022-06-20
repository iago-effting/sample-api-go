package main

import (
	"fmt"
	"iago-effting/api-example/pkg/http"
	"iago-effting/api-example/pkg/version"
	"log"
	"os"
)

func main() {
	fmt.Println("Version: ", version.Version)
	fmt.Println("Time Release:", version.Time)

	err := http.Run(":2020")

	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
