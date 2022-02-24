/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/harakeishi/whris/cmd"
)

var (
	version = "dev"
)

func main() {
	fmt.Print(version)
	cmd.Execute()
}
