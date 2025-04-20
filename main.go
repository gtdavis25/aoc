package main

import "github.com/alecthomas/kong"

var cli struct{}

func main() {
	context := kong.Parse(&cli)
	context.Run()
}
