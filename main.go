/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/harakeishi/whris/cmd"
)

var (
	version = "dev"
)

func main() {
	cmd.Execute(version)
}
