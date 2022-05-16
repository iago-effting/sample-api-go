package main

import (
	"fmt"
	"iago-effting/go-template/pkg/version"
)

func main() {
	fmt.Println("Version: ", version.Version)
	fmt.Println("Time Release:", version.Time)
}