package main

import (
	"cloudprobe/internal/g"
	"cloudprobe/internal/readfile"
	"cloudprobe/web"
)

func main() {

	//init config
	g.ParseConfig()

	// read file context
	readfile.NewServer().Do()

	//web
	web.Start()

}
