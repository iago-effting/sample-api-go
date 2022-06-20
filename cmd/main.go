package main

import (
	"fmt"
	"iago-effting/api-example/configs"
	"iago-effting/api-example/pkg/http"
	"iago-effting/api-example/pkg/version"
	"log"
	"os"
)

func main() {
	configs.SetVarEnvs(os.Getenv("ENV"))

	fmt.Println("Env: ", configs.Env.Name)
	fmt.Println("Version: ", version.Version)
	fmt.Println("Time Release:", version.Time)

	port := fmt.Sprintf(":%d", configs.Env.Server.Port)
	err := http.Run(port)

	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
